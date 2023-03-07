package models

// ErrorCode indicates the type of error
type ErrorCode string

const (
	// NotFound indicates that the resource is not found
	NotFound          ErrorCode = "not_found"
	// ServerError indicates that the server encountered an error
	ServerError       ErrorCode = "server_error"
	// TypeMismatch indicates that the type of the value is not correct
	TypeMismatch      ErrorCode = "type_mismatch"
	// ConflictError indicates that the resource already exists
	ConflictError     ErrorCode = "conflict"
)

// Response is the model for the response
type Response struct {
	StatusCode 	int    `json:"status_code"`
	Error 			string `json:"error"`
}