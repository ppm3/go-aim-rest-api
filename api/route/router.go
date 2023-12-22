package route

import (
	"context"
	"ppm3/go-aim-rest-api/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(ctx context.Context, c controllers.Controllers, router *gin.Engine) *gin.Engine {
	// Setup routes - Welcome
	router.GET("/", func(g *gin.Context) {
		g.JSON(200, gin.H{
			"message": "Welcome to the API",
		})
	})

	router.GET("/ping", c.PingController.Pong)
	router.GET("/health-check/api", c.HealthController.CheckHealth)
	router.GET("/health-check/db/mongo", c.HealthMongoDBController.CheckHealthDB)
	router.GET("/health-check/db/mysql", c.HealthMySQLController.CheckHealthDB)
	router.GET("/health-check/rabbitmq", c.HealthRabbitMQController.CheckHealthRabbitMQ)
	router.GET("/health-check/redis", c.HealthRedisController.CheckHealthRedis)

	return router
}
