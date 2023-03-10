package errors

import (
	"fmt"
	"net/http"
)

// RecordAlreadyExistsError is a custom error type for when a record already exists
type RecordAlreadyExistsError struct {
	Model string
}

func (e *RecordAlreadyExistsError) Error() string {
	return fmt.Sprintf("%s is already in the database.", e.Model)
}

// StatusCode returns the http status code for the error
func (e *RecordAlreadyExistsError) StatusCode() int {
	return http.StatusNotFound
}
