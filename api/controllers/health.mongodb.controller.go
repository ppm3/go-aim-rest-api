package controllers

import (
	"context"
	"net/http"
	"ppm3/go-aim-rest-api/configs"
	"ppm3/go-aim-rest-api/pkg/databases"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthMongoDBController struct {
	ctx          context.Context
	clients      *databases.Clients
	mongoActions databases.MongoDBActionsI
	configParams *configs.ServerConfig
}

func NewHealthMongoDBController(ctx context.Context, c *databases.Clients, m databases.MongoDBActionsI, p configs.ServerConfig) *HealthMongoDBController {
	return &HealthMongoDBController{
		clients:      c,
		mongoActions: m,
		configParams: &p,
		ctx:          ctx,
	}
}

func (h *HealthMongoDBController) Ping(c *gin.Context) {
	mc := h.clients.MongoDB

	uptime := time.Now().Unix()

	// Check the connection
	_, err := h.mongoActions.Ping(mc)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		_, err := h.mongoActions.PingDB(mc)
		if err != nil {
			c.AbortWithStatus(http.StatusBadGateway)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"uptime": uptime,
			})
		}
	}
}
