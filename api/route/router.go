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
	router.GET("/health-check/api", c.HealthController.Ping)
	router.GET("/health-check/db/mongo", c.HealthMongoDBController.Ping)
	router.GET("/health-check/db/mysql", c.HealthMySQLController.Ping)
	router.GET("/health-check/rabbitmq", c.HealthRabbitMQController.Ping)
	router.GET("/health-check/redis", c.HealthRedisController.Ping)

	return router
}
