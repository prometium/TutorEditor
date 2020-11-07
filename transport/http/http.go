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

	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("assets/images"))))

	r.Methods("GET").Path("/setup").Handler(httptransport.NewServer(
		e.SetupEndpoint,
		decodeSetupRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/script/raw").Handler(httptransport.NewServer(
		e.AddRawScriptEndpoint,
		decodeAddRawScriptRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/scripts").Handler(httptransport.NewServer(
		e.GetScriptsListEndpoint,
		decodeGetScriptsListRequest,
		encodeResponse,
	))

	return r
}

func decodeSetupRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req transport.SetupRequest
	return req, nil
}

func decodeAddRawScriptRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	scriptArchive, _, err := r.FormFile("script")
	if err != nil && err != http.ErrMissingFile {
		return nil, err
	}
	defer scriptArchive.Close()

	name := r.FormValue("name")

	return transport.AddRawScriptRequest{
		FileReader: scriptArchive,
		Name:       name,
	}, nil
}

func decodeGetScriptsListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req transport.GetScriptsListRequest
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
