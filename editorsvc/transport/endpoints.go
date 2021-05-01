package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/prometium/tutoreditor/editorsvc"
)

// Endpoints are exposed
type Endpoints struct {
	AddScriptArchiveEndpoint endpoint.Endpoint
	GetScriptArchiveEndpoint endpoint.Endpoint
	GetScriptsListEndpoint   endpoint.Endpoint
	GetScriptEndpoint        endpoint.Endpoint
	DeleteScriptEndpoint     endpoint.Endpoint
	UpdateScriptEndpoint     endpoint.Endpoint
	CopyScriptEndpoint       endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service
func MakeServerEndpoints(s editorsvc.Service) Endpoints {
	return Endpoints{
		AddScriptArchiveEndpoint: makeAddScriptArchiveEndpoint(s),
		GetScriptArchiveEndpoint: makeGetScriptArchiveEndpoint(s),
		GetScriptsListEndpoint:   makeGetScriptsListEndpoint(s),
		GetScriptEndpoint:        makeGetScriptEndpoint(s),
		DeleteScriptEndpoint:     makeDeleteScriptEndpoint(s),
		UpdateScriptEndpoint:     makeUpdateScriptEndpoint(s),
		CopyScriptEndpoint:       makeCopyScriptEndpoint(s),
	}
}

func makeAddScriptArchiveEndpoint(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddScriptArchiveRequest)
		uid, err := s.AddScriptArchive(ctx, req.Name, req.FileReader)
		return AddScriptArchiveResponse{UID: uid, Err: err}, nil
	}
}

func makeGetScriptArchiveEndpoint(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetScriptArchiveRequest)
		archiveBytes, err := s.GetScriptArchive(ctx, req.UID)
		return archiveBytes, err
	}
}

func makeGetScriptsListEndpoint(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		scripts, err := s.GetScriptsList(ctx)
		return GetScriptsListResponse{Scripts: scripts, Err: err}, nil
	}
}

func makeGetScriptEndpoint(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetScriptRequest)
		script, err := s.GetScript(ctx, req.UID)
		return GetScriptResponse{Script: script, Err: err}, nil
	}
}

func makeDeleteScriptEndpoint(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteScriptRequest)
		err := s.DeleteScript(ctx, req.UID)
		return DeleteScriptResponse{Err: err}, nil
	}
}

func makeUpdateScriptEndpoint(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateScriptRequest)
		uids, err := s.UpdateScript(ctx, req.Script, req.FrameIdsToDel, req.ActionIdsToDel)
		return UpdateScriptResponse{Uids: uids, Err: err}, nil
	}
}

func makeCopyScriptEndpoint(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CopyScriptRequest)
		uid, err := s.CopyScript(ctx, req.Script)
		return CopyScriptResponse{UID: uid, Err: err}, nil
	}
}
