package api

// ErrorCode indicates the type of error
type ErrorCode string

const (
	// NotFound indicates that the resource is not found
	NotFound          ErrorCode = "%s not found"
	// ServerError indicates that the server encountered an error
	ServerError       ErrorCode = "something went wrong!"
	// TypeMismatch indicates that the type of the value is not correct
	TypeMismatch      ErrorCode = "there is a type mismatch"
	// ConflictError indicates that the resource already exists
	ConflictError     ErrorCode = "conflict"
)

// Response is the model for the response
type Response struct {
	Message			string `json:"message,omitempty"`
}