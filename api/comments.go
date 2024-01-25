// api/comments.go

package api

import (
	"net/http"

	"github.com/anonymousa10/forum-backendplsda/models"
	"github.com/gin-gonic/gin"
)

// UpdateComment updates a comment by ID
// It expects a JSON body with the comment fields to be updated
// It returns the updated comment as JSON
func UpdateComment(c *gin.Context) {
	// Use the authenticated username from AuthMiddleware
	username, _ := c.Get("username")

	id := c.Param("id")

	var comment models.Comment
	// Use First to find a comment with the given ID
	result := models.DB.First(&comment, id)

	// Use Error to handle the error if the comment is not found
	if result.Error != nil {
		c.Error(result.Error)
		return
	}

	// Check if the authenticated user is the author of the comment
	if comment.Author != username {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this comment"})
		return
	}

	// Use BindJSON to bind the JSON body to the comment and abort if validation fails
	if err := c.BindJSON(&comment); err != nil {
		return
	}

	// Save the comment to the database
	models.DB.Save(&comment)

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// DeleteComment deletes a comment by ID
// It returns a success message as JSON
func DeleteComment(c *gin.Context) {
	// Use the authenticated username from AuthMiddleware
	username, _ := c.Get("username")

	id := c.Param("id")

	var comment models.Comment
	// Use First to find a comment with the given ID
	result := models.DB.First(&comment, id)

	// Use Error to handle the error if the comment is not found
	if result.Error != nil {
		c.Error(result.Error)
		return
	}

	// Check if the authenticated user is the author of the comment
	if comment.Author != username {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to delete this comment"})
		return
	}

	// Delete the comment from the database
	models.DB.Delete(&comment)

	// Use StatusText to get the standard text for the status code
	c.JSON(http.StatusOK, gin.H{"message": http.StatusText(http.StatusOK)})
}

// CreateComment creates a new comment
// It expects a JSON body with the comment fields
// It returns the created comment as JSON
func CreateComment(c *gin.Context) {
	// Use the authenticated username from AuthMiddleware
	username, _ := c.Get("username")

	var comment models.Comment
	// Use BindJSON to bind the JSON body to the comment and abort if validation fails
	if err := c.BindJSON(&comment); err != nil {
		return
	}

	// Set the comment author to the authenticated username
	comment.Author = username.(string)

	// Use FirstOrCreate to create a comment with the given fields or return the existing one
	models.DB.FirstOrCreate(&comment, comment)

	c.JSON(http.StatusCreated, gin.H{"data": comment})
}

// GetComment gets a comment by ID
// It returns the comment as JSON
func GetComment(c *gin.Context) {
	id := c.Param("id")

	var comment models.Comment
	// Use First to find a comment with the given ID
	result := models.DB.First(&comment, id)

	// Use Error to handle the error if the comment is not found
	if result.Error != nil {
		c.Error(result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// GetCommentsByThread gets all comments by thread ID
// It returns the comments and the authenticated username as JSON
func GetCommentsByThread(c *gin.Context) {
	// Use the authenticated username from AuthMiddleware
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	threadID := c.Param("threadID")

	var comments []models.Comment
	// Use Where and Find to get all comments with the given thread ID from the database
	result := models.DB.Where("thread_id = ?", threadID).Find(&comments)

	// Use Error to handle the error if the query fails
	if result.Error != nil {
		c.Error(result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comments, "username": username})
}
