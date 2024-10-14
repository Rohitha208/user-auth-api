package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProtectedHandler(c *gin.Context) {
	// This handler is protected; only authenticated users can access it
	c.JSON(http.StatusOK, gin.H{"message": "You are authorized to access this route!"})
}
