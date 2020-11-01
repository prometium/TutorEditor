package transport

type (
	// StatusRequest holds the request parameters for the Status method
	StatusRequest struct{}

	// StatusResponse holds the request parameters for the Status method
	StatusResponse struct {
		Status string `json:"status"`
	}

	// TransformScriptRequest holds the request parameters for the Status method
	TransformScriptRequest struct{
		Script string `json:"script"`
	}

	// TransformScriptResponse holds the request parameters for the Status method
	TransformScriptResponse struct {
		ID int `json:"id"`
	}
)
