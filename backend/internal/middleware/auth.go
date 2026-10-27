package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// In a real AWS environment, fetch this from AWS Secrets Manager
var jwtSecret = []byte("super_secret_mobile_app_key_change_me_later")

// RequireAuth is the security checkpoint for protected routes
func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Extract the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort() // Stops the request from proceeding
			return
		}

		// 2. Ensure it follows the industry-standard "Bearer <token>" format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 3. Parse and cryptographically verify the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verify the signing algorithm is what we expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// 4. Extract the user ID from the token payload and pass it to the next function
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// JWT stores numbers as float64, so we convert it back to a uint
			userID := uint(claims["user_id"].(float64))
			
			// Store the user ID in the Gin context so the final API handler knows who is making the request
			c.Set("user_id", userID)
			
			c.Next() // Request is authorized, allow it to continue to the API route
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to process token claims"})
			c.Abort()
			return
		}
	}
}