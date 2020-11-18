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
	AddBranchEndpoint      endpoint.Endpoint
	DeleteBranchEndpoint   endpoint.Endpoint
	DeleteFrameEndpoint    endpoint.Endpoint
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
		AddBranchEndpoint:      makeAddBranchEndpoint(s),
		DeleteBranchEndpoint:   makeDeleteBranchEndpoint(s),
		DeleteFrameEndpoint:    makeDeleteFrame(s),
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
		uids, err := s.UpdateScript(ctx, req.ID, req.Script)
		return UpdateScriptResponse{Uids: uids, Err: err}, nil
	}
}

func makeCopyScriptEndpoint(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CopyScriptRequest)
		id, err := s.CopyScript(ctx, req.Script)
		return CopyScriptResponse{ID: id, Err: err}, nil
	}
}

func makeAddBranchEndpoint(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddBranchRequest)
		uids, err := s.AddBranch(ctx, req.Branch)
		return AddBranchResponse{Uids: uids, Err: err}, nil
	}
}

func makeDeleteBranchEndpoint(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteBranchRequest)
		err := s.DeleteBranch(ctx, req.BranchToDelete)
		return DeleteBranchResponse{Err: err}, nil
	}
}

func makeDeleteFrame(s editorsvc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteFrameRequest)
		err := s.DeleteFrame(ctx, req.ID)
		return DeleteFrameResponse{Err: err}, nil
	}
}
