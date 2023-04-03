package controllers

import (
	"github.com/GorrillaProcess/go-crud/initializers"
	"github.com/GorrillaProcess/go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	initializers.LoadEnvVars()
	initializers.ConnectToDb()
	//Get data of req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	//Create post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// get id off url
	id := c.Param("id")

	//get the posts
	var post []models.Post
	initializers.DB.First(&post, id)

	//respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get the id off url
	id := c.Param("id")
	// Get the data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	// Find the post we update
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title,
		Body: body.Body})

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// Get the id off url
	id := c.Param("id")

	//Delete post
	initializers.DB.Delete(&models.Post{}, id)

	//Respond it
	c.Status(200)
}
