package controller

import (
	"finalproj/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type postInput struct {
	Title      string `json:"title" `
	Body       string `json:"body" `
	CategoryId uint   `json:"category_id"`
	UserId     uint   `json:"user_id"`
}

// GetAllPost godoc
// @Summary Get all Post.
// @Description Get a list of Post.
// @Tags Post
// @Produce json
// @Success 200 {object} []models.Post
// @Router /posts [get]
func GetAllPost(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var posts []models.Post
	db.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// GetPostById godoc
// @Summary Get Post.
// @Description Get an Post by id.
// @Tags Post
// @Produce json
// @Param id path string true "Post id"
// @Success 200 {object} models.Post
// @Router /posts/{id} [get]
func GetPostById(c *gin.Context) {
	var post models.Post

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "post tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// CreatePost godoc
// @Summary Create New Post.
// @Description Creating a new Post.
// @Tags Post
// @Param Body body postInput true "the body to create a new Post"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Post
// @Router /posts [post]
func CreatePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// validate Input
	var input postInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create post
	newPost := models.Post{Title: input.Title, Body: input.Body, UserId: input.UserId, CategoryId: input.CategoryId}
	db.Create(&newPost)

	c.JSON(http.StatusOK, gin.H{"data": newPost})
}

// UpdatePost godoc
// @Summary Update Post.
// @Description Update Post by id.
// @Tags Post
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Post id"
// @Param Body body postInput true "the body to update post"
// @Success 200 {object} map[string]boolean
// @Router /post/{id} [patch]
func UpdatePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// get model if exist
	var post models.Post
	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "post not found with such id"})
		return
	}

	// validate input
	var input postInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if post.UserId != input.UserId {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	var updatedPost models.Post
	updatedPost.Title = input.Title
	updatedPost.Body = input.Body

	db.Model(&post).Updates(&updatedPost)
	c.JSON(http.StatusOK, gin.H{"data": "update success"})
}

// DeletePost godoc
// @Summary Delete one Post.
// @Description Delete a Post by id.
// @Tags Post
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Post id"
// @Success 200 {object} map[string]boolean
// @Router /posts/{id} [delete]
func DeletePost(c *gin.Context) {
	// get model
	db := c.MustGet("db").(*gorm.DB)
	var post models.Post
	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "post not found with such id"})
		return
	}

	db.Delete(&post)

	c.JSON(http.StatusOK, gin.H{"message": "post deleted"})
}
