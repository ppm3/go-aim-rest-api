package controllers

import (
	"context"
	"net/http"
	"ppm3/go-aim-rest-api/configs"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	ctx          context.Context
	configParams *configs.ServerConfig
}

func NewHealthController(ctx context.Context, p configs.ServerConfig) *HealthController {
	return &HealthController{
		ctx:          ctx,
		configParams: &p,
	}
}

func (hc *HealthController) CheckHealth(c *gin.Context) {
	uptime := time.Now().Unix()
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"uptime":  uptime,
		"message": "API is up and running",
	})
}
