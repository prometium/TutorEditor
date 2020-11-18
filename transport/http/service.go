package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/prometium/tutoreditor/editorsvc"

	"github.com/go-kit/kit/log"
	kittransport "github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/prometium/tutoreditor/editorsvc/transport"
)

// MakeHTTPHandler mounts all of the service endpoints into an http.Handler
func MakeHTTPHandler(e transport.Endpoints, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("assets/images"))))

	r.Methods("POST").Path("/raw").Handler(httptransport.NewServer(
		e.AddRawScriptEndpoint,
		decodeAddRawScriptRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/scripts").Handler(httptransport.NewServer(
		e.GetScriptsListEndpoint,
		decodeGetScriptsListRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/scripts/{id}").Handler(httptransport.NewServer(
		e.GetScriptEndpoint,
		decodeGetScriptRequest,
		encodeResponse,
		options...,
	))

	r.Methods("DELETE").Path("/scripts/{id}").Handler(httptransport.NewServer(
		e.DeleteScriptEndpoint,
		decodeDeleteScriptRequest,
		encodeResponse,
		options...,
	))

	r.Methods("PUT").Path("/scripts/{id}").Handler(httptransport.NewServer(
		e.UpdateScriptEndpoint,
		decodeUpdateScriptRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/scripts").Handler(httptransport.NewServer(
		e.CopyScriptEndpoint,
		decodeCopyScriptRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/branches").Handler(httptransport.NewServer(
		e.AddBranchEndpoint,
		decodeAddBranchRequest,
		encodeResponse,
		options...,
	))

	r.Methods("DELETE").Path("/branches").Handler(httptransport.NewServer(
		e.DeleteBranchEndpoint,
		decodeDeleteBranchRequest,
		encodeResponse,
		options...,
	))

	r.Methods("DELETE").Path("/frames/{id}").Handler(httptransport.NewServer(
		e.DeleteFrameEndpoint,
		decodeDeleteFrameRequest,
		encodeResponse,
		options...,
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
	return transport.GetScriptRequest{ID: vars["id"]}, nil
}

func decodeDeleteScriptRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return transport.DeleteScriptRequest{ID: vars["id"]}, nil
}

func decodeUpdateScriptRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	var req transport.UpdateScriptRequest
	req.ID = vars["id"]
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeCopyScriptRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req transport.CopyScriptRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeAddBranchRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req transport.AddBranchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteBranchRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req transport.DeleteBranchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeDeleteFrameRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return transport.DeleteFrameRequest{ID: vars["id"]}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.Error() != nil {
		encodeError(ctx, e.Error(), w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	Error() error
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case editorsvc.ErrScriptNotFound:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
