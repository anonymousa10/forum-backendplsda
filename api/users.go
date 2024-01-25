// api/users.go

package api

import (
	"net/http"

	"github.com/anonymousa10/forum-backendplsda/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the username already exists
	if isUsernameTaken(user.Username) {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
		return
	}

	// Check if the email already exists
	if isEmailTaken(user.Email) {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already taken"})
		return
	}

	// Create the user if both username and email are unique
	models.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func isUsernameTaken(username string) bool {
	var existingUser models.User
	result := models.DB.Where("username = ?", username).First(&existingUser)
	return result.RowsAffected > 0
}

func isEmailTaken(email string) bool {
	var existingUser models.User
	result := models.DB.Where("email = ?", email).First(&existingUser)
	return result.RowsAffected > 0
}
func GetUser(c *gin.Context) {
	username := c.Param("username")

	var user models.User
	result := models.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
