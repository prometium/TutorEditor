package transport

import "io"

type (
	// StatusRequest holds the request parameters for the Status method
	StatusRequest struct{}

	// StatusResponse holds the request parameters for the Status method
	StatusResponse struct {
		Status string `json:"status"`
	}

	// AddRawScriptRequest holds the request parameters for the Status method
	AddRawScriptRequest struct {
		ArchiveReader io.Reader
	}

	// AddRawScriptResponse holds the request parameters for the Status method
	AddRawScriptResponse struct {
		ID int `json:"id"`
	}
)
