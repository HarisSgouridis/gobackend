package repository

import (
	"github.com/HarisSgouridis/gobackend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRepository struct {
	router *gin.Engine
}

func NewUserRepository() *UserRepository {
	repo := &UserRepository{}
	return repo
}

func (repo *UserRepository) createUser(c *gin.Context) {
	// Parse the request and create a new user
	var newUser model.User // Assuming you have a User struct defined
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Perform user creation logic, e.g., saving to a database
	// ...

	// Return a success response
	c.JSON(http.StatusCreated, newUser)
}
