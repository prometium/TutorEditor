package transport

import (
	"editorsvc"
	"io"
)

// AddRawScriptRequest holds the request parameters for the AddRawScript method
type AddRawScriptRequest struct {
	Name       string `json:"name"`
	FileReader io.ReadCloser
}

// AddRawScriptResponse holds the response parameters for the AddRawScript method
type AddRawScriptResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error,omitempty"`
}

func (r AddRawScriptResponse) Error() error { return r.Err }

// GetScriptsListRequest holds the request parameters for the GetScriptsList method
type GetScriptsListRequest struct{}

// GetScriptsListResponse holds the response parameters for the GetScriptsList method
type GetScriptsListResponse struct {
	Scripts []editorsvc.Script `json:"scripts"`
	Err     error              `json:"error,omitempty"`
}

func (r GetScriptsListResponse) Error() error { return r.Err }

// GetScriptRequest holds the request parameters for the GetScript method
type GetScriptRequest struct {
	ID string `json:"id"`
}

// GetScriptResponse holds the response parameters for the GetScript method
type GetScriptResponse struct {
	Script *editorsvc.Script `json:"script"`
	Err    error             `json:"error,omitempty"`
}

func (r GetScriptResponse) Error() error { return r.Err }

// DeleteScriptRequest holds the request parameters for the DeleteScript method
type DeleteScriptRequest struct {
	ID string `json:"id"`
}

// DeleteScriptResponse holds the response parameters for the DeleteScript method
type DeleteScriptResponse struct {
	Err error `json:"error,omitempty"`
}

func (r DeleteScriptResponse) Error() error { return r.Err }

// UpdateScriptRequest holds the request parameters for the UpdateScript method
type UpdateScriptRequest struct {
	Script *editorsvc.Script `json:"script"`
}

// UpdateScriptResponse holds the response parameters for the UpdateScript method
type UpdateScriptResponse struct {
	Err error `json:"error,omitempty"`
}

func (r UpdateScriptResponse) Error() error { return r.Err }

// AddBranchPointRequest holds the request parameters for the AddBranchPoint method
type AddBranchPointRequest struct {
	BranchPoint *editorsvc.BranchPoint `json:"branchPoint"`
}

// AddBranchPointResponse holds the response parameters for the AddBranchPoint method
type AddBranchPointResponse struct {
	Uids map[string]string `json:"uids"`
	Err  error             `json:"error,omitempty"`
}

func (r AddBranchPointResponse) Error() error { return r.Err }
