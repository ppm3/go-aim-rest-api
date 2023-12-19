package controllers

import (
	"context"
	"net/http"
	"ppm3/go-aim-rest-api/configs"

	"github.com/gin-gonic/gin"
)

type PingController struct {
	ctx          context.Context
	configParams *configs.ServerConfig
}

func NewPingController(ctx context.Context, p configs.ServerConfig) *PingController {
	return &PingController{
		ctx:          ctx,
		configParams: &p,
	}
}

func (hc *PingController) Pong(c *gin.Context) {
	c.String(http.StatusOK, "PONG")
}
