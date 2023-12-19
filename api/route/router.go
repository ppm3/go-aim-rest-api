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

	router.GET("/health-check", c.HealthController.CheckHealth)
	router.GET("/ping", c.PingController.Pong)

	return router
}
