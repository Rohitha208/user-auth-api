package handlers

import (
	"net/http"

	"user-auth-api/services"
	"user-auth-api/src/api/dto"

	"github.com/gin-gonic/gin"
)

func LoginHandler(loginService *services.LoginService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequest dto.LoginRequest
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		token, err := loginService.Login(loginRequest.Username, loginRequest.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
