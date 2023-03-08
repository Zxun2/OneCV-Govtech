package errors

import "net/http"

// BadRequestError is a custom error type for when a request is bad
type BadRequestError struct {
	Message string
}

func (e *BadRequestError) Error() string {
	if e.Message == "" {
		return "Bad Request"
	}
	return e.Message
}

// StatusCode returns the http status code for the error
func (e *BadRequestError) StatusCode() int {
	return http.StatusBadRequest
}