// api/auth.go

package api

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/anonymousa10/forum-backendplsda/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("secretkey") // Replace with a secure secret key

// AuthMiddleware is a middleware to handle JWT authentication
// AuthMiddleware is a middleware to handle JWT authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		var tokenString string
		if authHeader != "" {
			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) == 2 {
				tokenString = bearerToken[1]
			}
		}

		// If no token was found in the Authorization header, try to get it from the cookies
		if tokenString == "" {
			tokenCookie, _ := c.Cookie("jwt_token")
			if tokenCookie != "" {
				tokenString = tokenCookie
			}
		}

		if tokenString == "" {
			fmt.Println("Missing Authorization header or cookie")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			fmt.Println("Error parsing JWT token:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok || !token.Valid {
			fmt.Println("Invalid JWT token or claims:", ok, token.Valid)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		fmt.Println("Authenticated user:", claims.Subject)

		c.Set("username", claims.Subject)

		c.Next() // Continue to the next handler
	}
}

// Login handles user login and issues a JWT token
func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the entered username and password are valid
	if isValidUser(&user) {
		// Create a JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
			Subject:   user.Username,
		})

		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			fmt.Printf("Error creating JWT token: %s\n", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create JWT token"})
			return
		}

		// Set the token as a cookie
		c.SetCookie("jwt_token", tokenString, int((time.Hour * 24).Seconds()), "/", "", false, true)

		c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenString})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	}
}

// Logout handles user logout by clearing the JWT token cookie
func Logout(c *gin.Context) {
	c.SetCookie("jwt_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

// isValidUser checks if the entered username and password are valid
func isValidUser(user *models.User) bool {
	var existingUser models.User
	result := models.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&existingUser)

	return result.RowsAffected > 0
}
