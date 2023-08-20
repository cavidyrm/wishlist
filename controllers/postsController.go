package controllers

import (
	"crud/initializers"
	"crud/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data of request
	var body struct {
		Body  string
		Title string
	}
	if err := c.Bind(&body); err != nil {
		fmt.Println(err)
	}
	post := models.Post{Title: body.Title, Body: body.Body}

	// Create a post
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	// Return data
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts data
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with posts

	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostsShow(c *gin.Context) {
	//get id from url
	id := c.Param("id")
	//find the post
	var post models.Post
	initializers.DB.First(&post, id)
	//get the specified post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")
	// Get the data off request body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)
	// Update the post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	// Respond the post
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsDelete(c *gin.Context) {
	// Get the id from url
	id := c.Param("id")
	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)
	// Respond it
	c.Status(200)
}
