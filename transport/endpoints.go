package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"editorsvc"
)

// Endpoints are exposed
type Endpoints struct {
	StatusEndpoint       endpoint.Endpoint
	AddRawScriptEndpoint endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service
func MakeServerEndpoints(s editorsvc.Service) Endpoints {
	return Endpoints{
		StatusEndpoint:       makeStatusEndpoint(s),
		AddRawScriptEndpoint: makeAddRawScriptEndpoint(s),
	}
}

func makeStatusEndpoint(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(StatusRequest)
		status, err := s.Status(ctx)
		if err != nil {
			return StatusResponse{status}, err
		}

		return StatusResponse{status}, nil
	}
}

func makeAddRawScriptEndpoint(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRawScriptRequest)
		id, err := s.AddRawScript(ctx, req.ArchiveReader)
		if err != nil {
			return AddRawScriptResponse{id}, err
		}

		return AddRawScriptResponse{id}, nil
	}
}
