// main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/anonymousa10/forum-backendplsda/api"
	"github.com/anonymousa10/forum-backendplsda/models"

	// Import cors package
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// Initialize the database connection
	models.InitDB()

	// Create a new Gin router
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	// Define non-authenticated routes
	router.POST("/api/login", api.Login)
	router.POST("/api/logout", api.Logout)
	router.POST("/api/signup", api.CreateUser)
	router.GET("/api/users/:username", api.GetUser)

	// Apply AuthMiddleware to the following routes
	authGroup := router.Group("/api")
	authGroup.Use(api.AuthMiddleware())
	{
		authGroup.GET("/threads/:id", api.GetThread)            // Move the more specific route first
		authGroup.GET("/threads/page/:page", api.GetAllThreads) // Updated route with pagination

		authGroup.POST("/threads", api.CreateThread)

		authGroup.PUT("/threads/:id", api.UpdateThread)
		authGroup.DELETE("/threads/:id", api.DeleteThread)

		authGroup.GET("/comments/thread/:threadID", api.GetCommentsByThread)
		authGroup.POST("/comments", api.CreateComment)
		authGroup.GET("/comments/:id", api.GetComment)
		authGroup.PUT("/comments/:id", api.UpdateComment)
		authGroup.DELETE("/comments/:id", api.DeleteComment)

		authGroup.GET("/categories", api.GetAllCategories)
		authGroup.POST("/categories", api.CreateCategory)
		authGroup.GET("/categories/:id", api.GetCategory)
		authGroup.PUT("/categories/:id", api.UpdateCategory)
		authGroup.DELETE("/categories/:id", api.DeleteCategory)

		// Add other authenticated routes here
	}

	// Start the server
	port := ":8080"
	fmt.Printf("Server is listening on port %s...\n", port)
	http.ListenAndServe(port, router) // Use the router directly
}
