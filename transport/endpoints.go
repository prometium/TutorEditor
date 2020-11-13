package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"editorsvc"
)

// Endpoints are exposed
type Endpoints struct {
	AddRawScriptEndpoint   endpoint.Endpoint
	GetScriptsListEndpoint endpoint.Endpoint
	GetScriptEndpoint      endpoint.Endpoint
	DeleteScriptEndpoint   endpoint.Endpoint
	UpdateScriptEndpoint   endpoint.Endpoint
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
	}
}

func makeAddRawScriptEndpoint(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRawScriptRequest)
		id, err := s.AddRawScript(ctx, req.Name, req.FileReader)
		return AddRawScriptResponse{ID: id, Err: err}, nil
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
		script, err := s.GetScript(ctx, req.ID)
		return GetScriptResponse{Script: script, Err: err}, nil
	}
}

func makeDeleteScriptEndpoint(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteScriptRequest)
		err := s.DeleteScript(ctx, req.ID)
		return DeleteScriptResponse{Err: err}, nil
	}
}

func makeUpdateScriptEndpoint(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateScriptRequest)
		err := s.UpdateScript(ctx, req.Script)
		return UpdateScriptResponse{Err: err}, nil
	}
}
