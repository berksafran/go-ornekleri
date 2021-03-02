package model

// ResponseSuccess is struct of success response as a JSON format.
type ResponseSuccess struct {
	ID      interface{} `json:"id"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Path    interface{} `json:"path,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ResponseError is struct of error response as a JSON format.
type ResponseError struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
