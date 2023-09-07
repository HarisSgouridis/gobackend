package controller

import (
	"fmt"
	"github.com/HarisSgouridis/gobackend/model"
	"github.com/HarisSgouridis/gobackend/mongo"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func InitializeRoutes(router *gin.Engine) {
	client, err := mongo.NewMongoDBClient(mongo.MongoDBConfig{
		URI:      "mongodb+srv://Haris:Theoharis@db2mongo.ddnmcb9.mongodb.net/?retryWrites=true&w=majority",
		Database: "bfs",
	})

	router.POST("/users", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err != nil {
			// Handle the error, e.g., log it or return an error response
			fmt.Println("Failed to create MongoDB client:", err)
			// You might want to return an error response to the client here.
			return
		}

		//defer func(client *mongo.MongoDBClient) {
		//	err := client.Close()
		//	if err != nil {
		//		panic("Error disconnecting from MongoDB: " + err.Error())
		//	}
		//}(client)

		// Insert the new user into MongoDB with the provided data
		err = client.CreateUser(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		// Return a success response with the newly created user
		c.JSON(http.StatusCreated, user)
	})

	router.GET("/getUser", func(c *gin.Context) {
		email := c.DefaultQuery("email", "")

		filter := bson.M{"email": email}

		user, err := client.ReadUser(&filter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// If the user is not found, return a custom error message
		if user == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		// Return the user data in the response
		c.JSON(http.StatusOK, user)
	})

	router.PUT("/updateUser", func(c *gin.Context) {

		email := c.DefaultQuery("email", "")

		filter := bson.M{"email": email}

		var user2 model.User
		if err := c.ShouldBindJSON(&user2); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		update := bson.M{"$set": bson.M{
			"username": user2.UserName,
			"password": user2.PassWord,
			"email":    user2.Email,
		}}

		err := client.UpdateUser(filter, update)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

	})

	router.DELETE("/deleteUser", func(c *gin.Context) {

		email := c.DefaultQuery("email", "")

		filter := bson.M{"email": email}

		err := client.DeleteUser(filter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

	})
}
