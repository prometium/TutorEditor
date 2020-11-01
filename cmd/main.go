package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"

	"editorsvc"
	"editorsvc/dgraphdb"
	"editorsvc/implementation"
	"editorsvc/transport"
	httptransport "editorsvc/transport/http"
)

func main() {
	var httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var dgraphClient *dgo.Dgraph
	{
		conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		defer conn.Close()

		dgraphClient = dgo.NewDgraphClient(api.NewDgraphClient(conn))
	}

	var service editorsvc.Service
	{
		repository, err := dgraphdb.New(dgraphClient)
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
		handler = httptransport.MakeHTTPHandler(endpoints)
	}

	errs := make(chan error)

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	logger.Log("exit", <-errs)
}
