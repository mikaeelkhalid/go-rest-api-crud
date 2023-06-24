package controllers

import (
	"example.com/m/v2/initializers"
	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
)


func CreatePosts(c *gin.Context) {

	// get req body
	var body struct {
		Title string
		Body string
	}

	c.Bind(&body)

	//create
	post := models.Post{Title: body.Title, Body:body.Body}

    result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {

	var posts []models.Post 

	initializers.DB.Find(&posts)

	// return it
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {

	id := c.Param("id")

	var post models.Post 

	initializers.DB.First(&post, id)

	// return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {

	id := c.Param("id")

	// get req body
	var body struct {
		Title string
		Body string
	}

	c.Bind(&body)

	var post models.Post 
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// return it
	c.JSON(200, gin.H{
		"post": post,
	})
}