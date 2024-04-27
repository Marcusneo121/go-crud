package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/marcus121neo/go-crud/initializers"
	"github.com/marcus121neo/go-crud/models"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostByID(c *gin.Context) {
	// Get ID off URL
	id := c.Param("id")

	// Get the posts
	var post models.Post
	initializers.DB.Find(&post, id)

	if post.ID == 0 {
		c.JSON(400, gin.H{
			"error":   "error",
			"message": "Post not found",
		})
		return
	}

	// Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get ID off URL
	id := c.Param("id")

	// Get the data of request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	// Find the post were updating
	var post models.Post
	initializers.DB.Find(&post, id)

	if post.ID == 0 {
		c.JSON(400, gin.H{
			"error":   "error",
			"message": "Unable to update, Post not found",
		})
		return
	}

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// Get ID off URL
	id := c.Param("id")

	// Find the post were updating
	var post models.Post
	initializers.DB.Find(&post, id)

	if post.ID == 0 {
		c.JSON(400, gin.H{
			"error":   "error",
			"message": "Unable to delete, Post not found",
		})
		return
	}

	// Update it
	// To permanantly delete use Unscoped
	//db.Unscoped().Delete(&Books, 3)

	initializers.DB.Delete(&post, id)

	// Respond with it
	c.JSON(200, gin.H{
		"message": "Post Deleted. Post ID: " + id,
	})
}
