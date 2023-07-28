package controller

import (
	"finalproj/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type commentInput struct {
	Content string `json:"content"`
	// PostId string `json:"post_id"`
	UserId uint `json:"user_id"`
}

// GetAllCommentByPostId godoc
// @Summary Get all post Comment.
// @Description Get a list of post Comment.
// @Tags Comment
// @Produce json
// @Param id path string true "Post id"
// @Success 200 {object} []models.Comment
// @Router /posts/{id}/comments [get]
func GetcommentsByPostId(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	postId := c.Param("id")

	var comments []models.Comment
	if err := db.Where("post_id = ?", postId).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "tidak ada komentar pada postingan"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// GetCommentById godoc
// @Summary Get Comment.
// @Description Get a Comment by id.
// @Tags Comment
// @Produce json
// @Param id path string true "Comment id"
// @Success 200 {object} models.Comment
// @Router /comments/{id} [get]
func GetCommentById(c *gin.Context) {
	var comment models.Comment
	var commentId = c.Param("id")

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", commentId).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "comment tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// CreateComment godoc
// @Summary Create New Comment.
// @Description Creating a new Comment.
// @Tags Comment
// @Param Body body commentInput true "the body to create a new Comment"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Comment id"
// @Success 200 {object} models.Comment
// @Router /posts/{id}/comments [post]
func CreateComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	postId, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate input
	var input commentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create comment
	newComment := models.Comment{Content: input.Content, UserID: input.UserId, PostID: postId}
	db.Create(&newComment)

	c.JSON(http.StatusOK, gin.H{"data": newComment})
}

// UpdateComment godoc
// @Summary Update Comment.
// @Description Update Comment by id.
// @Tags Comment
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Comment id"
// @Param Body body commentInput true "the body to update comment"
// @Success 200 {object} map[string]boolean
// @Router /comments/{id} [patch]
func UpdateComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// get model if exist
	var comment models.Comment
	var commentId = c.Param("id")
	if err := db.Where("id = ?", commentId).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "comment not found with such id"})
		return
	}

	// validate input
	var input commentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedComment models.Comment
	updatedComment.Content = input.Content

	db.Model(&comment).Updates(&updatedComment)
	c.JSON(http.StatusOK, gin.H{"data": "update success"})
}

// DeleteComment godoc
// @Summary Delete one Comment.
// @Description Delete a Comment by id.
// @Tags Comment
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Comment id"
// @Success 200 {object} map[string]boolean
// @Router /comments/{id} [delete]
func DeleteComment(c *gin.Context) {
	// get model
	db := c.MustGet("db").(*gorm.DB)
	var comment models.Comment
	var commentId = c.Param("id")
	if err := db.Where("id = ? ", commentId).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "comment not found"})
		return
	}

	db.Delete(&comment)
	c.JSON(http.StatusOK, gin.H{"message": "comment deleted"})

}
