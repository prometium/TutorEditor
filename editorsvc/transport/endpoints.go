package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/prometium/tutoreditor/editorsvc"
)

// Endpoints are exposed
type Endpoints struct {
	AddRawScriptEndpoint   endpoint.Endpoint
	GetScriptsListEndpoint endpoint.Endpoint
	GetScriptEndpoint      endpoint.Endpoint
	DeleteScriptEndpoint   endpoint.Endpoint
	UpdateScriptEndpoint   endpoint.Endpoint
	CopyScriptEndpoint     endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service
func MakeServerEndpoints(s editorsvc.Service) Endpoints {
	return Endpoints{
		AddRawScriptEndpoint:   makeAddRawScriptEndpoint(s),
		GetScriptsListEndpoint: makeGetScriptsListEndpoint(s),
		GetScriptEndpoint:      makeGetScriptEndpoint(s),
		DeleteScriptEndpoint:   makeDeleteScriptEndpoint(s),
		UpdateScriptEndpoint:   makeUpdateScriptEndpoint(s),
		CopyScriptEndpoint:     makeCopyScriptEndpoint(s),
	}
}

func makeAddRawScriptEndpoint(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRawScriptRequest)
		uid, err := s.AddRawScript(ctx, req.Name, req.FileReader)
		return AddRawScriptResponse{UID: uid, Err: err}, nil
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
