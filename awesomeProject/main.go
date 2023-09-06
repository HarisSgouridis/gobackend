package main

import (
	"github.com/HarisSgouridis/gobackend/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize routes
	controller.InitializeRoutes(userRepo.Router)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Go and Gin!")
	})
	r.Run()

}
