// api/categories.go

package api

import (
	"net/http"

	"github.com/anonymousa10/forum-backendplsda/models"
	"github.com/gin-gonic/gin"
)

// UpdateCategory updates a category by ID
// It expects a JSON body with the category fields to be updated
// It returns the updated category as JSON
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	var category models.Category
	// Use FirstOrInit to find or initialize a category with the given ID
	models.DB.FirstOrInit(&category, id)

	// Use BindJSON to bind the JSON body to the category and abort if validation fails
	if err := c.BindJSON(&category); err != nil {
		return
	}

	// Save the category to the database
	models.DB.Save(&category)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// DeleteCategory deletes a category by ID
// It returns a success message as JSON
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	var category models.Category
	// Use First to find a category with the given ID
	result := models.DB.First(&category, id)

	// Use Error to handle the error if the category is not found
	if result.Error != nil {
		c.Error(result.Error)
		return
	}

	// Delete the category from the database
	models.DB.Delete(&category)

	// Use StatusText to get the standard text for the status code
	c.JSON(http.StatusOK, gin.H{"message": http.StatusText(http.StatusOK)})
}

// CreateCategory creates a new category
// It expects a JSON body with the category fields
// It returns the created category as JSON
func CreateCategory(c *gin.Context) {
	var category models.Category
	// Use BindJSON to bind the JSON body to the category and abort if validation fails
	if err := c.BindJSON(&category); err != nil {
		return
	}

	// Use FirstOrCreate to create a category with the given fields or return the existing one
	models.DB.FirstOrCreate(&category, category)

	c.JSON(http.StatusCreated, gin.H{"data": category})
}

// GetCategory gets a category by ID
// It returns the category as JSON
func GetCategory(c *gin.Context) {
	id := c.Param("id")

	var category models.Category
	// Use First to find a category with the given ID
	result := models.DB.First(&category, id)

	// Use Error to handle the error if the category is not found
	if result.Error != nil {
		c.Error(result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// GetAllCategories gets all categories
// It returns the categories as JSON
func GetAllCategories(c *gin.Context) {
	var categories []models.Category
	// Use Find to get all categories from the database
	result := models.DB.Find(&categories)

	// Use Error to handle the error if the query fails
	if result.Error != nil {
		c.Error(result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}
