package errors

import (
	"fmt"
	"net/http"
)

// RecordNotFoundError is a custom error type for when a record is not found
type RecordNotFoundError struct {
	Model string
}

func (e *RecordNotFoundError) Error() string {
	return fmt.Sprintf("%s could not be found.", e.Model)
}

// StatusCode returns the http status code for the error
func (e *RecordNotFoundError) StatusCode() int {
	return http.StatusNotFound
}
