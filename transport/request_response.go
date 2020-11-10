package transport

import (
	"editorsvc"
	"io"
)

type (
	// AddRawScriptRequest holds the request parameters for the AddRawScript method
	AddRawScriptRequest struct {
		Name       string `json:"name"`
		FileReader io.ReadCloser
	}

	// AddRawScriptResponse holds the response parameters for the AddRawScript method
	AddRawScriptResponse struct {
		ID string `json:"id"`
	}

	// GetScriptsListRequest holds the request parameters for the GetScriptsList method
	GetScriptsListRequest struct{}

	// GetScriptsListResponse holds the response parameters for the GetScriptsList method
	GetScriptsListResponse struct {
		Scripts []editorsvc.Script `json:"scripts"`
	}

	// GetScriptRequest holds the request parameters for the GetScript method
	GetScriptRequest struct {
		ID string `json:"id"`
	}

	// GetScriptResponse holds the response parameters for the GetScript method
	GetScriptResponse struct {
		Script []editorsvc.Script `json:"script"`
	}
)
