package controllers

import (
	"context"
	"net/http"
	"ppm3/go-aim-rest-api/configs"
	"ppm3/go-aim-rest-api/pkg/databases"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthMySQLController struct {
	ctx          context.Context
	clients      *databases.Clients
	mysqlActions databases.MySQLActionsI
	configParams *configs.ServerConfig
}

func NewHealthMySQLController(ctx context.Context, c *databases.Clients, m databases.MySQLActionsI, p configs.ServerConfig) *HealthMySQLController {
	return &HealthMySQLController{
		clients:      c,
		mysqlActions: m,
		configParams: &p,

		ctx: ctx,
	}
}

func (h *HealthMySQLController) Ping(c *gin.Context) {
	mysqlCon := h.clients.MySQL
	uptime := time.Now().Unix()

	// Check the connection
	_, err := h.mysqlActions.Ping(mysqlCon)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"uptime": uptime,
		})
	}

}
