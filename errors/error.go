package errors

import (
	"Zxun2/OneCV-Govtech/api"
	"net/http"
)

// MakeResponseCode returns the http status code based on the error code
func MakeResponseCode(response api.Response) int {
	if len(response.Message) == 0 {
		return http.StatusOK
	}
	return http.StatusInternalServerError
}

// MakeResponseErr returns the response with the error code
func MakeResponseErr(err error) api.Response {
	return api.Response{Message: err.Error()}
}