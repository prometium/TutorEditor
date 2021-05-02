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

	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/",
		cacheControlWrapper(
			http.FileServer(
				http.Dir("assets/images"),
			),
		),
	))

	r.Methods("POST").Path("/archive").Handler(httptransport.NewServer(
		e.AddScriptArchiveEndpoint,
		decodeAddScriptArchiveRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/archiveV2").Handler(httptransport.NewServer(
		e.AddScriptArchiveV2Endpoint,
		decodeAddScriptArchiveRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/archiveV2/{uid}").Handler(httptransport.NewServer(
		e.GetScriptArchiveEndpoint,
		decodeGetScriptArchiveRequest,
		encodeBytesResponse,
		options...,
	))

	r.Methods("GET").Path("/scripts").Handler(httptransport.NewServer(
		e.GetScriptsListEndpoint,
		decodeGetScriptsListRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/scripts/{uid}").Handler(httptransport.NewServer(
		e.GetScriptEndpoint,
		decodeGetScriptRequest,
		encodeResponse,
		options...,
	))

	r.Methods("DELETE").Path("/scripts/{uid}").Handler(httptransport.NewServer(
		e.DeleteScriptEndpoint,
		decodeDeleteScriptRequest,
		encodeResponse,
		options...,
	))

	r.Methods("PUT").Path("/scripts").Handler(httptransport.NewServer(
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

	return r
}

func cacheControlWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age=0, must-revalidate")
		h.ServeHTTP(w, r)
	})
}

func decodeAddScriptArchiveRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	file, _, err := r.FormFile("script")
	if err != nil && err != http.ErrMissingFile {
		return nil, err
	}
	name := r.FormValue("name")
	return transport.AddScriptArchiveRequest{
		FileReader: file,
		Name:       name,
	}, nil
}

func decodeGetScriptArchiveRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return transport.GetScriptArchiveRequest{UID: vars["uid"]}, nil
}

func decodeGetScriptsListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req transport.GetScriptsListRequest
	return req, nil
}

func decodeGetScriptRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return transport.GetScriptRequest{UID: vars["uid"]}, nil
}

func decodeDeleteScriptRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return transport.DeleteScriptRequest{UID: vars["uid"]}, nil
}

func decodeUpdateScriptRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req transport.UpdateScriptRequest
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

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	if err, ok := response.(errorer); ok && err.Error() != nil {
		encodeError(ctx, err.Error(), w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeBytesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/zip")
	w.Write(response.([]byte))
	return nil
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
	case editorsvc.ErrFileNotAttached:
		return http.StatusBadRequest
	case editorsvc.ErrScriptNotFound:
		return http.StatusBadRequest
	case editorsvc.ErrVersionsDoNotMatch:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
