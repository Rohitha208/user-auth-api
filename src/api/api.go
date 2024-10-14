package api

import (
	"user-auth-api/services"
	"user-auth-api/src/api/handlers"
	"user-auth-api/src/api/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the router and sets up the routes
func SetupRouter(loginService *services.LoginService, registrationService *services.RegistrationService) *gin.Engine {
	router := gin.Default()

	// Register routes with the provided services
	router.POST("/register", handlers.RegisterHandler(registrationService))
	router.POST("/login", handlers.LoginHandler(loginService))
	router.GET("/protected", middleware.ValidToken(), handlers.ProtectedHandler)

	return router
}
