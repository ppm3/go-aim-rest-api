package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"ppm3/go-aim-rest-api/configs"

	"github.com/gin-gonic/gin"
)

// Load configuration
var projectDirName string = "go-aim-rest-api"

func main() {
	var configParams *configs.ServerConfig
	var err error

	ctx := context.Background()

	configParams, err = configs.Load(ctx, os.Getenv("ENVIRONMENT"), projectDirName)
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Create a new Gin-gonic router
	router := gin.Default()

	// Set up routes
	setupRoutes(router)

	// Start the server
	log.Printf("Server is running on port %s", configParams.Server.Port)
	err = router.Run(":" + configParams.Server.Port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func setupRoutes(router *gin.Engine) {
	// Set up your routes here
	// For example:
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
}
