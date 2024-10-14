package middleware

import (
	"net/http"
	"strings"

	"user-auth-api/src/api/helpers" // Adjust based on your helpers package

	"github.com/gin-gonic/gin"
)

func ValidToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not provided"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(token, "Bearer ")

		// Validate the token (using a method you have defined in your helpers)
		if _, err := helpers.ValidateToken(tokenString); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// If the token is valid, proceed to the next handler
		c.Next()
	}
}
