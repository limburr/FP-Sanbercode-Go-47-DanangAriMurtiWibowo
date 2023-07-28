package controller

import (
	"finalproj/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type categoryInput struct {
	Name string `json:"name"`
}

// GetAllCategory godoc
// @Summary Get all Category.
// @Description Get a list of Category.
// @Tags Category
// @Produce json
// @Success 200 {object} []models.Category
// @Router /categories [get]
func GetAllCategory(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var categories []models.Category
	db.Find(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// CreateCategory godoc
// @Summary Create New Category.
// @Description Creating a new Category.
// @Tags Category
// @Param Body body categoryInput true "the body to create a new Category"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Category
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	// Validate input
	var input categoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Rating
	category := models.Category{Name: input.Name}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&category)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// GetCategoryById godoc
// @Summary Get Category.
// @Description Get an Category by id.
// @Tags Category
// @Produce json
// @Param id path string true "Category id"
// @Success 200 {object} models.Category
// @Router /categories/{id} [get]
func GetCategoryById(c *gin.Context) {
	var category models.Category

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// UpdateCategory godoc
// @Summary Update Category.
// @Description Update Category by id.
// @Tags Category
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Category id"
// @Param Body body categoryInput true "the body to update post"
// @Success 200 {object} models.Category
// @Router /categories/{id} [patch]
func UpdateCategory(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var category models.Category

	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input categoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Category
	updatedInput.Name = input.Name
	updatedInput.UpdatedAt = time.Now()

	db.Model(&category).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// DeleteCategory godoc
// @Summary Delete one Category.
// @Description Delete a Category by id.
// @Tags Category
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "Category id"
// @Success 200 {object} map[string]boolean
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var category models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&category)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
