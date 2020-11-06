package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"editorsvc"
)

// Endpoints are exposed
type Endpoints struct {
	SetupEndpoint          endpoint.Endpoint
	AddRawScriptEndpoint   endpoint.Endpoint
	GetScriptsListEndpoint endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service
func MakeServerEndpoints(s editorsvc.Service) Endpoints {
	return Endpoints{
		SetupEndpoint:          makeSetupEndpoint(s),
		AddRawScriptEndpoint:   makeAddRawScriptEndpoint(s),
		GetScriptsListEndpoint: makeGetScriptsListEndpoint(s),
	}
}

func makeSetupEndpoint(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		err := s.Setup(ctx)
		return SetupResponse{}, err
	}
}

func makeAddRawScriptEndpoint(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRawScriptRequest)
		id, err := s.AddRawScript(ctx, req.Name, req.FileReader)
		return AddRawScriptResponse{id}, err
	}
}

func makeGetScriptsListEndpoint(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		scripts, err := s.GetScriptsList(ctx)
		return GetScriptsListResponse{scripts}, err
	}
}
