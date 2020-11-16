package transport

import (
	"github.com/prometium/tutoreditor/editorsvc"
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
	ID     string            `json:"id"`
	Script *editorsvc.Script `json:"script"`
}

// UpdateScriptResponse holds the response parameters for the UpdateScript method
type UpdateScriptResponse struct {
	Uids map[string]string `json:"uids"`
	Err  error             `json:"error,omitempty"`
}

func (r UpdateScriptResponse) Error() error { return r.Err }

// CopyScriptRequest holds the request parameters for the UpdateScript method
type CopyScriptRequest struct {
	Script *editorsvc.Script `json:"script"`
}

// CopyScriptResponse holds the response parameters for the CopyScript method
type CopyScriptResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error,omitempty"`
}

func (r CopyScriptResponse) Error() error { return r.Err }

// AddBranchRequest holds the request parameters for the AddBranch method
type AddBranchRequest struct {
	Branch *editorsvc.Branch `json:"branch"`
}

// AddBranchResponse holds the response parameters for the AddBranch method
type AddBranchResponse struct {
	Uids map[string]string `json:"uids"`
	Err  error             `json:"error,omitempty"`
}

func (r AddBranchResponse) Error() error { return r.Err }

// DeleteBranchRequest holds the request parameters for the DeleteBranch method
type DeleteBranchRequest struct {
	BranchToDelete *editorsvc.BranchToDelete `json:"branchToDelete"`
}

// DeleteBranchResponse holds the response parameters for the DeleteBranch method
type DeleteBranchResponse struct {
	Err error `json:"error,omitempty"`
}

func (r DeleteBranchResponse) Error() error { return r.Err }
