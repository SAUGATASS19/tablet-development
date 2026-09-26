package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Initialize the Gin engine with default settings (includes logging and crash recovery)
	r := gin.Default()

	// 2. Define a route using a GET request
	r.GET("/", func(c *gin.Context) {
		// Send back a JSON response instead of plain text
		c.JSON(200, gin.H{
			"message": "Hello from the Gin backend!",
			"status":  "success",
		})
	})

	// 3. Start the server on port 8080
	r.Run(":8080")
}
