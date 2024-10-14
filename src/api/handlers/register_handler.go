package handlers

import (
	"net/http"

	"user-auth-api/services"
	"user-auth-api/src/api/dto"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(registrationService *services.RegistrationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var registerRequest dto.RegisterUserRequest
		if err := c.ShouldBindJSON(&registerRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		err := registrationService.Register(registerRequest) // Adjust based on your service method
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
	}
}
