package controller

import (
	"github.com/HarisSgouridis/gobackend/model"
	"github.com/HarisSgouridis/gobackend/mongo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitializeRoutes(router *gin.Engine) {
	// Create a new user
	router.POST("/users", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Insert the new user into MongoDB with the provided data
		err := mongo.MongoDBClient.CreateUser(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		// Return a success response with the newly created user
		c.JSON(http.StatusCreated, user)
	})

	// Add other routes for updating, deleting, listing users, etc.
}
