package controllers

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	initializers "github.com/saurabhranjan007/go-server/Initializers"
	models "github.com/saurabhranjan007/go-server/Models"
)

// #### HOME ####
func Home(c *gin.Context) {

	server_port := os.Getenv("PORT")

	c.JSON(200, gin.H{
		"message": "Server is running at: " + server_port,
	})
}

// #### CREATE POST ####
func PostCreate(c *gin.Context) {

	// Get the request body and bind 
	var req_body struct {
		Body  string
		Title string
	}
	c.Bind(&req_body)

	// Req data to insert 
	post := models.Post{
		Title: req_body.Title,
		Body:  req_body.Body,
	}

	result := initializers.DB.Create(&post)

	// Check & Return for Error
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

// #### GET ALL POSTS ####
func ReadPosts(c *gin.Context) {

	var posts []models.Post

	initializers.DB.Find(&posts)

	// Return posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// #### GET POST(:id) - Single ####
func GetPost(c *gin.Context) {

	// Get ID form URL
	id := c.Param("id")

	var post models.Post

	initializers.DB.Find(&post, id)

	// Return post
	c.JSON(200, gin.H{
		"posts": post,
	})
}

// #### UPDATE POST ####
func PostUpdate(c *gin.Context) {

	fmt.Println("PUT Method Called ..")

	// Get the id from URL
	id := c.Param("id")

	// Get data off the request body (bind as well)
	var req_body struct {
		Body  string
		Title string
	}
	c.Bind(&req_body)

	// Find the data to update
	var post models.Post
	initializers.DB.Find(&post, id)

	// Update Data
	initializers.DB.Model(&post).Updates(models.Post{
		Title: req_body.Title,
		Body:  req_body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

// #### DELETE POST ####
func DeletePost(c *gin.Context) {

	// Get ID form URL
	id := c.Param("id")

	initializers.DB.Delete(&models.Post{}, id)

	c.Status(200)
	return
}
