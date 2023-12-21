package main

import (
	"context"
	"log"
	"os"
	"ppm3/go-aim-rest-api/api/controllers"
	"ppm3/go-aim-rest-api/api/route"
	"ppm3/go-aim-rest-api/configs"
	"ppm3/go-aim-rest-api/pkg/databases"
	"ppm3/go-aim-rest-api/pkg/rabbitmq"

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

	// Load configuration
	var ctx context.Context = context.Background()

	configParams, err := config.Load(ctx, os.Getenv("ENVIRONMENT"), projectDirName)
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Set Gin-gonic mode
	gin.SetMode(configParams.Server.Mode)

	// Create a new app
	var app App = *newApp(
		ctx,
		*configParams,
	)

	// Connect to MongoDB
	var mongoDBActions databases.MongoDBActionsI = databases.NewMongoDBActions(app.ctx, &configParams.Mongo)

	mongoClient, err := mongoDBActions.MongoConnect()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Connect to MySQL
	var mysqlDBActions databases.MySQLActionsI = databases.NewMySQLActions(app.ctx, &configParams.Mysql)

	mysqlClient, err := mysqlDBActions.MySQLConnect()
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}

	// Create a new clients
	var clientsDB databases.Clients = *databases.NewClients(mongoClient, mysqlClient)

	// // RabbitMQ connection
	var rabbitMQActions rabbitmq.RabbitMQActionsI = rabbitmq.NewRabbitMQConnect(app.ctx, &configParams.RabbitMQ)
	rabbitClient, err := rabbitMQActions.Connect()
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}

	// Create a new controllers
	var controllers controllers.Controllers = *controllers.NewControllers(
		ctx,
		&clientsDB,
		rabbitClient,
		mongoDBActions,
		mysqlDBActions,
		rabbitMQActions,
		*configParams,
	)

	// Create a new Gin-gonic router
	app.Router = route.SetupRouter(app.ctx, controllers, app.Router)

	// Start the server
	log.Printf("Server is running on port %s", configParams.Server.Port)
	err = app.Router.Run(":" + configParams.Server.Port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
