package transport

import (
	"io"

	"github.com/prometium/tutoreditor/editorsvc"
)

// AddScriptArchiveRequest holds the request parameters for the AddScriptArchive method
type AddScriptArchiveRequest struct {
	Name       string `json:"name"`
	FileReader io.ReadCloser
}

// AddScriptArchiveResponse holds the response parameters for the AddScriptArchive method
type AddScriptArchiveResponse struct {
	UID string `json:"uid"`
	Err error  `json:"error,omitempty"`
}

func (r AddScriptArchiveResponse) Error() error { return r.Err }

// AddScriptArchiveRequest holds the request parameters for the AddScriptArchive method
type GetScriptArchiveRequest struct {
	UID string `json:"uid"`
}

// ReleaseScriptArchiveRequest holds the request parameters for the ReleaseScript method
type ReleaseScriptArchiveRequest struct {
	UID string `json:"uid"`
}

// ReleaseScriptArchiveResponse holds the response parameters for the ReleaseScript method
type ReleaseScriptArchiveResponse struct {
	Err error `json:"error,omitempty"`
}

func (r ReleaseScriptArchiveResponse) Error() error { return r.Err }

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
	UID string `json:"uid"`
}

// GetScriptResponse holds the response parameters for the GetScript method
type GetScriptResponse struct {
	Script *editorsvc.Script `json:"script"`
	Err    error             `json:"error,omitempty"`
}

func (r GetScriptResponse) Error() error { return r.Err }

// DeleteScriptRequest holds the request parameters for the DeleteScript method
type DeleteScriptRequest struct {
	UID string `json:"uid"`
}

// DeleteScriptResponse holds the response parameters for the DeleteScript method
type DeleteScriptResponse struct {
	Err error `json:"error,omitempty"`
}

func (r DeleteScriptResponse) Error() error { return r.Err }

// UpdateScriptRequest holds the request parameters for the UpdateScript method
type UpdateScriptRequest struct {
	UID            string            `json:"uid"`
	Script         *editorsvc.Script `json:"script"`
	FrameIdsToDel  []string          `json:"frameIdsToDel"`
	ActionIdsToDel []string          `json:"actionIdsToDel"`
}

// UpdateScriptResponse holds the response parameters for the UpdateScript method
type UpdateScriptResponse struct {
	Uids map[string]string `json:"uids"`
	Err  error             `json:"error,omitempty"`
}

func (r UpdateScriptResponse) Error() error { return r.Err }

// CopyScriptRequest holds the request parameters for the CopyScript method
type CopyScriptRequest struct {
	Script *editorsvc.Script `json:"script"`
}

// CopyScriptResponse holds the response parameters for the CopyScript method
type CopyScriptResponse struct {
	UID string `json:"uid"`
	Err error  `json:"error,omitempty"`
}

func (r CopyScriptResponse) Error() error { return r.Err }

// AddImageRequest holds the request parameters for the AddImage method
type AddImageRequest struct {
	FileReader io.ReadCloser
}

// AddImageResponse holds the response parameters for the AddImage method
type AddImageResponse struct {
	Link string `json:"link,omitempty"`
	Err  error  `json:"error,omitempty"`
}

func (r AddImageResponse) Error() error { return r.Err }
