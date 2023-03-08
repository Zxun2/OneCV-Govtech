package controllers

import (
	"Zxun2/OneCV-Govtech/errors"
	"Zxun2/OneCV-Govtech/models"
	"Zxun2/OneCV-Govtech/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Suspend suspends a student
func Suspend(c *gin.Context) {
	var payload models.SuspendStudentPayload

	err := c.ShouldBindJSON(&payload)
	if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errors.MakeResponseErr(err),
		)
	}
	response := services.SuspendStudent(payload)
	c.JSON(errors.MakeResponseCode(response.Response), response)
}