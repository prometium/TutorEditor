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

	r.Methods("POST").Path("/raw").Handler(httptransport.NewServer(
		e.AddRawScriptEndpoint,
		decodeAddRawScriptRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/scripts").Handler(httptransport.NewServer(
		e.GetScriptsListEndpoint,
		decodeGetScriptsListRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/script/{id}").Handler(httptransport.NewServer(
		e.GetScriptEndpoint,
		decodeGetScriptRequest,
		encodeResponse,
	))

	r.Methods("DELETE").Path("/script/{id}").Handler(httptransport.NewServer(
		e.DeleteScriptEndpoint,
		decodeDeleteScriptRequest,
		encodeResponse,
	))

	return r
}

func decodeAddRawScriptRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	file, _, err := r.FormFile("script")
	if err != nil && err != http.ErrMissingFile {
		return nil, err
	}
	name := r.FormValue("name")
	return transport.AddRawScriptRequest{
		FileReader: file,
		Name:       name,
	}, nil
}

func decodeGetScriptsListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req transport.GetScriptsListRequest
	return req, nil
}

func decodeGetScriptRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return transport.GetScriptRequest{
		ID: vars["id"],
	}, nil
}

func decodeDeleteScriptRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return transport.DeleteScriptRequest{
		ID: vars["id"],
	}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
