package main

import (
	"example.com/m/v2/controllers"
	"example.com/m/v2/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabase()

}

func main() {
	route := gin.Default()

	route.POST("/create", controllers.CreatePosts)
	route.GET("/posts", controllers.GetPosts)
	route.PUT("/post/:id", controllers.UpdatePost)
	route.GET("/post/:id", controllers.GetPost)
	
	
	route.Run() // listen and serve on 0.0.0.0:8080
}