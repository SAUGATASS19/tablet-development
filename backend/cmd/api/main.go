package main

import (
	"backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Initialize the Gin engine with default settings (includes logging and crash recovery)
	r := gin.Default()

	// 2. Register the health check route
	r.GET("/health", handlers.HealthCheck)

	// 3. Start the server on port 8080
	r.Run(":8080")
}
