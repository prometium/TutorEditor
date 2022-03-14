package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/rs/cors"
	"google.golang.org/grpc"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/prometium/tutoreditor/editorsvc"
	dgraphdb "github.com/prometium/tutoreditor/editorsvc/database"
	"github.com/prometium/tutoreditor/editorsvc/implementation"
	"github.com/prometium/tutoreditor/editorsvc/transport"
	httptransport "github.com/prometium/tutoreditor/editorsvc/transport/http"
	"github.com/prometium/tutoreditor/editorsvc/utils"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	err := godotenv.Load()
	if err != nil {
		level.Info(logger).Log("Error loading .env file")
	}

	var httpAddr = flag.String("http.addr", fmt.Sprintf("%s:%s", utils.Getenv("APP_HOST", ""), utils.Getenv("APP_PORT", "9001")), "HTTP listen address")
	flag.Parse()

	var ctx = context.Background()

	var dgraphClient *dgo.Dgraph
	{
		conn, err := grpc.Dial(fmt.Sprintf("%s:%s", utils.Getenv("DB_HOST", "db-alpha"), utils.Getenv("DB_PORT", "9080")), grpc.WithInsecure())
		defer conn.Close()
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

		dgraphClient = dgo.NewDgraphClient(api.NewDgraphClient(conn))
	}

	var minioClient *minio.Client
	{
		accessKeyID := utils.Getenv("S3_ACCESS_KEY_ID", "minioadmin")
		secretAccessKey := utils.Getenv("S3_SECRET_ACCESS_KEY", "minioadmin")
		useSSL := false
		minioClient, err = minio.New(fmt.Sprintf("%s:%s", utils.Getenv("S3_HOST", "s3"), utils.Getenv("S3_PORT", "9099")), &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: useSSL,
		})
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

		location := utils.Getenv("S3_LOCATION", "us-east-1")

		bucketName := utils.Getenv("S3_BUCKET_NAME", "editor")
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
			if errBucketExists == nil && exists {
				level.Info(logger).Log(fmt.Sprintf("Backet with name %s already exists\n", bucketName))
			} else {
				level.Error(logger).Log("exit", err)
				os.Exit(-1)
			}
		} else {
			level.Info(logger).Log(fmt.Sprintf("Backet with name %s successfully created \n", bucketName))
		}
		policy := fmt.Sprintf(`{
			"Version": "2012-10-17",
			"Statement": [
			  	{
					"Effect": "Allow",
					"Principal": { "AWS": ["*"] },
					"Action": ["s3:GetObject"],
					"Resource": ["arn:aws:s3:::%s/*"],
					"Sid": ""
			  	}
			]
		}`, bucketName)
		err = minioClient.SetBucketPolicy(ctx, bucketName, policy)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

		sharedBucketName := utils.Getenv("S3_SHARED_BUCKET_NAME", "archive")
		err = minioClient.MakeBucket(ctx, sharedBucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			exists, errBucketExists := minioClient.BucketExists(ctx, sharedBucketName)
			if errBucketExists == nil && exists {
				level.Info(logger).Log(fmt.Sprintf("Backet with name %s already exists\n", sharedBucketName))
			} else {
				level.Error(logger).Log("exit", err)
				os.Exit(-1)
			}
		} else {
			level.Info(logger).Log(fmt.Sprintf("Backet with name %s successfully created \n", sharedBucketName))
		}
		policy = fmt.Sprintf(`{
			"Version": "2012-10-17",
			"Statement": [
			  	{
					"Effect": "Allow",
					"Principal": { "AWS": ["*"] },
					"Action": ["s3:GetObject"],
					"Resource": ["arn:aws:s3:::%s/*"],
					"Sid": ""
			  	}
			]
		}`, sharedBucketName)
		err = minioClient.SetBucketPolicy(ctx, sharedBucketName, policy)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	var service editorsvc.Service
	{
		repository := dgraphdb.New(dgraphClient)
		err := repository.Setup(ctx)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

		service = implementation.NewService(repository, minioClient)
	}

	var endpoints transport.Endpoints
	{
		endpoints = transport.MakeServerEndpoints(service)
	}

	var handler http.Handler
	{
		router := httptransport.MakeHTTPHandler(endpoints, logger)
		cors := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedHeaders:   []string{"X-Requested-With"},
			AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"},
			AllowCredentials: true,
		})

		handler = cors.Handler(router)
	}

	errs := make(chan error)

	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", *httpAddr)
		server := &http.Server{
			Addr:    *httpAddr,
			Handler: handler,
		}
		errs <- server.ListenAndServe()
	}()

	rand.Seed(time.Now().UnixNano())

	level.Error(logger).Log("exit", <-errs)
}
