package models

type ErrorResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

type SuccessResponse struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}
