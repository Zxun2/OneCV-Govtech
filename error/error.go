package error

import (
	"Zxun2/OneCV-Govtech/models"
	"net/http"
)

// MakeResponseCode returns the http status code based on the error code
func MakeResponseCode(response models.Response) int {
	if len(response.Error) == 0 {
		return http.StatusOK
	}
	return http.StatusInternalServerError
}

// MakeResponseErr returns the response with the error code
func MakeResponseErr(err models.ErrorCode) models.Response {
	return models.Response{Error: string(err)}
}