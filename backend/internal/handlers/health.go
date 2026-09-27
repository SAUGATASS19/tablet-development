package handlers 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// HealthCheck responds to AWS and monitoring tools to confirm the server is alive
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "Server is running perfectly",
	})
}