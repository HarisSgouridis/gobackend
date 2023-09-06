package main

import (
	"github.com/HarisSgouridis/gobackend/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize routes

	r := gin.Default()

	controller.InitializeRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Go and Gin!")
	})
	r.Run()

}
