package http

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"editorsvc/transport"
)

// MakeHTTPHandler mounts all of the service endpoints into an http.Handler
func MakeHTTPHandler(e transport.Endpoints) http.Handler {
	r := mux.NewRouter()

	r.Methods("GET").Path("/status").Handler(httptransport.NewServer(
		e.StatusEndpoint,
		decodeStatusRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/script").Handler(httptransport.NewServer(
		e.AddRawScriptEndpoint,
		decodeAddRawScriptRequest,
		encodeResponse,
	))

	return r
}

func decodeStatusRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req transport.StatusRequest
	return req, nil
}

func decodeAddRawScriptRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	scriptArchive, _, err := r.FormFile("script")
	if err != nil && err != http.ErrMissingFile {
		return nil, err
	}
	defer scriptArchive.Close()

	return transport.AddRawScriptRequest{
		ArchiveReader: scriptArchive,
	}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
