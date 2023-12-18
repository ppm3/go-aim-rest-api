package main

import (
	"context"
	"log"
	"os"
	"ppm3/go-aim-rest-api/api/controllers"
	"ppm3/go-aim-rest-api/api/route"
	"ppm3/go-aim-rest-api/configs"

	"github.com/gin-gonic/gin"
)

var config *configs.LoadServerConfig

type App struct {
	Router       *gin.Engine
	ctx          context.Context
	configParams *configs.ServerConfig
}

// Load configuration
var projectDirName string = "go-aim-rest-api"

func newApp(ctx context.Context, p configs.ServerConfig) *App {
	return &App{
		ctx:          ctx,
		Router:       gin.Default(),
		configParams: &p,
	}
}

func main() {
	var configParams *configs.ServerConfig
	var err error

	var ctx context.Context = context.Background()

	configParams, err = config.Load(ctx, os.Getenv("ENVIRONMENT"), projectDirName)
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	var app App = *newApp(
		ctx,
		*configParams,
	)

	var controllers controllers.Controllers = *controllers.NewControllers(ctx, *configParams)

	// Create a new Gin-gonic router
	app.Router = route.SetupRouter(app.ctx, controllers, app.Router)

	// Start the server
	log.Printf("Server is running on port %s", configParams.Server.Port)
	err = app.Router.Run(":" + configParams.Server.Port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
