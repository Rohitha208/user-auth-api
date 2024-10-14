package main

import (
	"user-auth-api/services"
	"user-auth-api/src/api"
)

func main() {
	// Create your services
	registrationService := services.NewRegistrationService()
	loginService := services.NewLoginService()

	// Set up the router
	router := api.SetupRouter(loginService, registrationService)

	// Start the server
	router.Run(":8080") // listen and serve on 8080
}
