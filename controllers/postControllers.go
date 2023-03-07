package controllers

import (
	"GO-CRUD1/initializers"
	"GO-CRUD1/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
		Name  string
		Desc  string
	}
	c.Bind(&body)
	post := models.Post{Body: body.Body, Title: body.Title, Name: body.Name, Desc: body.Desc}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(500)
		return
	}
	c.JSON(200, post)
}
func UpdateDB(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Desc  string
		Body  string
		Title string
		Name  string
	}
	c.Bind(&body)
	var post models.Post
	initializers.DB.First(&post, id)
	initializers.DB.Model(&post).Updates(models.Post{
		Desc:  body.Desc,
		Body:  body.Body,
		Title: body.Title,
		Name:  body.Name,
	})
	c.JSON(200, post)
}
func ShowAll(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, posts)
}
func ShowByID(c *gin.Context) {
	id := c.Param("id")
	var posts models.Post
	initializers.DB.First(&posts, id)
	c.JSON(200, posts)
}
