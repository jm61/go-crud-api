package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jm61/restApi/initializers"
	"github.com/jm61/restApi/models"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostSingle(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	result := initializers.DB.First(&post, id)
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")
	result := initializers.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	c.Status(200)
}
