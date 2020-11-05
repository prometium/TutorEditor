package transport

import "io"

type (
	// SetupRequest holds the request parameters for the Setup method
	SetupRequest struct{}

	// SetupResponse holds the response parameters for the Setup method
	SetupResponse struct{}

	// AddRawScriptRequest holds the request parameters for the AddRawScript method
	AddRawScriptRequest struct {
		Name          string `json:"name"`
		ArchiveReader io.Reader
	}

	// AddRawScriptResponse holds the response parameters for the AddRawScript method
	AddRawScriptResponse struct {
		ID string `json:"id"`
	}
)
