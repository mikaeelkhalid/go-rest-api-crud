package main

import (
	"example.com/m/v2/initializers"
	"example.com/m/v2/models"
)

func init() {
	initializers.ConnectDatabase()
	initializers.LoadEnv()
}

func main(){
	initializers.DB.AutoMigrate(&models.Post{})
}