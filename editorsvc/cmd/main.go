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

	"github.com/prometium/tutoreditor/editorsvc"
	dgraphdb "github.com/prometium/tutoreditor/editorsvc/database"
	"github.com/prometium/tutoreditor/editorsvc/implementation"
	"github.com/prometium/tutoreditor/editorsvc/transport"
	httptransport "github.com/prometium/tutoreditor/editorsvc/transport/http"
	"github.com/prometium/tutoreditor/editorsvc/utils"
)

func main() {
	var httpAddr = flag.String("http.addr", fmt.Sprintf("0.0.0.0:%s", utils.Getenv("APP_PORT", "9000")), "HTTP listen address")
	flag.Parse()

	var ctx = context.Background()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var dgraphClient *dgo.Dgraph
	{
		conn, err := grpc.Dial(fmt.Sprintf("0.0.0.0:%s", utils.Getenv("DB_PORT", "9080")), grpc.WithInsecure())
		defer conn.Close()
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

		dgraphClient = dgo.NewDgraphClient(api.NewDgraphClient(conn))
	}

	var service editorsvc.Service
	{
		repository := dgraphdb.New(dgraphClient)
		err := repository.Setup(ctx)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

		service = implementation.NewService(repository)
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
