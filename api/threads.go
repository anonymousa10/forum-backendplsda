// api/threads.go

package api

import (
	"net/http"
	"strconv"

	"github.com/anonymousa10/forum-backendplsda/models"
	"github.com/gin-gonic/gin"
)

// UpdateThread updates a thread by ID
// It expects a JSON body with the thread fields to be updated
// It returns the updated thread as JSON
func UpdateThread(c *gin.Context) {
	// Use the authenticated username from AuthMiddleware
	username, _ := c.Get("username")

	id := c.Param("id")

	var thread models.Thread
	// Use FirstOrInit to find or initialize a thread with the given ID
	models.DB.FirstOrInit(&thread, id)

	// Use BindJSON to bind the JSON body to the thread and abort if validation fails
	if err := c.BindJSON(&thread); err != nil {
		return
	}

	// Check if the authenticated user is the author of the thread
	if thread.Author != username {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this thread"})
		return
	}

	// Save the thread to the database
	models.DB.Save(&thread)

	c.JSON(http.StatusOK, gin.H{"data": thread})
}

// DeleteThread deletes a thread by ID
// It returns a success message as JSON
func DeleteThread(c *gin.Context) {
	// Use the authenticated username from AuthMiddleware
	username, _ := c.Get("username")

	id := c.Param("id")

	var thread models.Thread
	// Use First to find a thread with the given ID
	result := models.DB.First(&thread, id)

	// Use Error to handle the error if the thread is not found
	if result.Error != nil {
		c.Error(result.Error)
		return
	}

	// Check if the authenticated user is the author of the thread
	if thread.Author != username {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to delete this thread"})
		return
	}

	// Delete the thread from the database
	models.DB.Delete(&thread)

	// Use StatusText to get the standard text for the status code
	c.JSON(http.StatusOK, gin.H{"message": http.StatusText(http.StatusOK)})
}

// CreateThread creates a new thread
// It expects a JSON body with the thread fields
// It returns the created thread as JSON
func CreateThread(c *gin.Context) {
	// Use the authenticated username from AuthMiddleware
	username, _ := c.Get("username")

	var thread models.Thread
	// Use BindJSON to bind the JSON body to the thread and abort if validation fails
	if err := c.BindJSON(&thread); err != nil {
		return
	}

	// Set the thread author to the authenticated username
	thread.Author = username.(string)

	// Use FirstOrCreate to create a thread with the given fields or return the existing one
	models.DB.FirstOrCreate(&thread, thread)

	c.JSON(http.StatusCreated, gin.H{"data": thread})
}

// GetThread gets a thread by ID
// It returns the thread as JSON
func GetThread(c *gin.Context) {
	id := c.Param("id")

	var thread models.Thread
	// Use First to find a thread with the given ID
	result := models.DB.First(&thread, id)

	// Use Error to handle the error if the thread is not found
	if result.Error != nil {
		c.Error(result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": thread})
}

// GetAllThreads gets all threads with optional offset
// It returns the threads as JSON
func GetAllThreads(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit := 10
	offset := (page - 1) * limit

	var threads []models.Thread
	// Use Offset and Limit to get the threads with pagination from the database
	result := models.DB.Offset(offset).Limit(limit).Find(&threads)

	// Use Error to handle the error if the query fails
	if result.Error != nil {
		c.Error(result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": threads})
}
