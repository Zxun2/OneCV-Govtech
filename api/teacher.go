package api

import (
	"Zxun2/OneCV-Govtech/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RetrieveNotifications retrieves a list of students who can receive a given notification.
func (s *Server) RetrieveNotifications(c *gin.Context) {
	var payload ListStudentReceivingNotificationPayload

	err := c.ShouldBindJSON(&payload)
	if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, makeResponseErr(err),
		)
	}
	students, err := services.ListStudentsReceiveNotifications(s.store, payload.Teacher, payload.Notification)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ListStudentReceivingNotificationResponse{
			Response: makeResponseErr(err),
		})
		return	
	}

	c.JSON(http.StatusOK, ListStudentReceivingNotificationResponse{
		Recipients: students,
	})
}

// Register multiple new students to a specified teacher
func (s *Server) Register(c *gin.Context) {
	var payload RegistersStudentsPayload

	err := c.ShouldBindJSON(&payload)
	if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, makeResponseErr(err),
		)
	}

	err = services.RegisterStudentsToTeacher(s.store, payload.Teacher, payload.Students)
	if err != nil {
		c.JSON(http.StatusInternalServerError, RegistersStudentsResponse{
			Response: makeResponseErr(err),
		})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}

// GetCommonStudents retrieves a list of students common to a given list of teachers
// The list of teachers is specified in the query parameters.
func (s *Server) GetCommonStudents(c *gin.Context) {
	teachers := c.QueryArray("teacher")
	log.Println(teachers)
	students, err := services.GetCommonStudents(s.store, teachers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, RetrieveCommonStudentsResponse{
			Response: makeResponseErr(err),
		})
		return
	}

	c.JSON(http.StatusOK, RetrieveCommonStudentsResponse{
		Students: students,
	})
}

