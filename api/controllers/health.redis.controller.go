package controllers

import (
	"context"
	"net/http"
	"ppm3/go-aim-rest-api/configs"
	redisAction "ppm3/go-aim-rest-api/pkg/redis"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type HealthRedisController struct {
	ctx          context.Context
	redisClient  *redis.Client
	configParams *configs.ServerConfig
	redisActions redisAction.RedisActionsI
}

func NewHealthRedisController(ctx context.Context, r *redis.Client, rA redisAction.RedisActionsI, p configs.ServerConfig) *HealthRedisController {
	return &HealthRedisController{
		ctx:          ctx,
		redisClient:  r,
		configParams: &p,
		redisActions: rA,
	}
}

func (h *HealthRedisController) Ping(c *gin.Context) {

	_, err := h.redisActions.Ping(h.redisClient)

	if err != nil {
		c.AbortWithStatus(http.StatusBadGateway)
	}

	uptime := time.Now().Unix()
	c.JSON(http.StatusOK, gin.H{
		"uptime": uptime,
	})
}
