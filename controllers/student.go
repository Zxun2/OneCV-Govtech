package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Suspend suspends a student
func Suspend(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}