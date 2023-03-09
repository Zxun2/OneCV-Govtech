package api

import (
	"Zxun2/OneCV-Govtech/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Suspend suspends a student
func (s *Server) Suspend(c *gin.Context) {
	var payload SuspendStudentPayload

	err := c.ShouldBindJSON(&payload)
	if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, makeResponseErr(err))
			return
	}
	err = services.SuspendStudent(s.store, payload.Student)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, makeResponseErr(err))
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}