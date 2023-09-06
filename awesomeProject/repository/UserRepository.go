package repository

import (
	"github.com/HarisSgouridis/gobackend/model"
	"github.com/HarisSgouridis/gobackend/mongo"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRepository struct {
	router *gin.Engine
}

func NewUserRepository() *UserRepository {
	repo := &UserRepository{}

	mongoConfig := mongo.MongoDBConfig{
		URI:      "mongodb+srv://Haris:Theoharis@db2mongo.ddnmcb9.mongodb.net/?retryWrites=true&w=majority", // Replace with your MongoDB URI
		Database: "bfs",                                                                                     // Replace with your database name
	}

	mongoClient, err := mongo.NewMongoDBClient(mongoConfig)

	if err != nil {
		panic(err) // Handle the error appropriately in your application
	}

	return repo
}

func (repo *UserRepository) CreateUser(c *gin.Context) {
	var newUser model.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the new user into MongoDB
	if err := mongo.MongoDBClient.CreateUser(newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Return a success response with the newly created user
	c.JSON(http.StatusCreated, newUser)
}
