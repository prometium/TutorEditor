package transport

import "io"

type (
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
