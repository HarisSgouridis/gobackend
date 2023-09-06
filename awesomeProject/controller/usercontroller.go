package controller

import (
	"github.com/HarisSgouridis/gobackend/repository"
	"github.com/gin-gonic/gin"
)

func () initializeRoutes() {
	router := gin.Default()

	// Create a new user
	router.POST("/users", repository.NewUserRepository().CreateUser)

	// Add other routes for updating, deleting, listing users, etc.
}
