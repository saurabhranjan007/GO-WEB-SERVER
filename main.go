package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	initializers "github.com/saurabhranjan007/go-server/Initializers"
	"github.com/saurabhranjan007/go-server/controllers"
)

// SPECIAL FUNC: init() - runs right before main function
func init() {
	fmt.Println("$$$$$ MAIN INITIALIZERS CALLED $$$$$")
	initializers.LoadEnvvariables()
	initializers.DatabaseConnection()
}

func main() {

	r := gin.Default()

	// API Endpoints
	r.GET("/", controllers.Home)
	r.POST("/post", controllers.PostCreate)
	r.GET("/posts", controllers.ReadPosts)
	r.GET("/post/:id", controllers.GetPost)
	r.PUT("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.DeletePost)

	r.Run()
}
