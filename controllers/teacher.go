package controllers

import (
	"Zxun2/OneCV-Govtech/errors"
	"Zxun2/OneCV-Govtech/models"
	"Zxun2/OneCV-Govtech/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RetrieveNotifications retrieves a list of students who can receive a given notification.
func RetrieveNotifications(c *gin.Context) {
	var payload models.ListStudentReceivingNotificationPayload

	err := c.ShouldBindJSON(&payload)
	if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errors.MakeResponseErr(err),
		)
	}
	response := services.ListStudentsReceiveNotifications(payload)
	c.JSON(http.StatusOK, response)
}

// Register multiple new students to a specified teacher
func Register(c *gin.Context) {
	var payload models.RegistersStudentsPayload

	err := c.ShouldBindJSON(&payload)
	if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errors.MakeResponseErr(err),
		)
	}

	response := services.RegisterStudentsToTeacher(payload)
	if response.Message != "" {
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

// GetCommonStudents retrieves a list of students common to a given list of teachers
// The list of teachers is specified in the query parameters.
func GetCommonStudents(c *gin.Context) {
	teachers := c.QueryArray("teacher")
	log.Println(teachers)
	students, err := services.GetCommonStudents(teachers)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.RetrieveCommonStudentsResponse{
			Response: errors.MakeResponseErr(err),
		})
		return
	}

	c.JSON(http.StatusOK, models.RetrieveCommonStudentsResponse{
		Students: students,
	})
}

