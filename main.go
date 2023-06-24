package main

import (
	"example.com/m/v2/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabase()

}

func main() {
	route := gin.Default()

	route.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	
	route.Run() // listen and serve on 0.0.0.0:8080
}