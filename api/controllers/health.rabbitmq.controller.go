package controllers

import (
	"context"
	"net/http"
	"ppm3/go-aim-rest-api/configs"
	"ppm3/go-aim-rest-api/pkg/rabbitmq"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type HealthRabbitMQController struct {
	ctx             context.Context
	rabbitConnect   *amqp.Connection
	configParams    *configs.ServerConfig
	rabbitMQActions rabbitmq.RabbitMQActionsI
}

func NewHealthRabbitMQController(ctx context.Context, r *amqp.Connection, rMQ rabbitmq.RabbitMQActionsI, p configs.ServerConfig) *HealthRabbitMQController {
	return &HealthRabbitMQController{
		ctx:             ctx,
		rabbitConnect:   r,
		configParams:    &p,
		rabbitMQActions: rMQ,
	}
}

func (h *HealthRabbitMQController) Ping(c *gin.Context) {
	var conn *amqp.Connection = h.rabbitConnect
	uptime := time.Now().Unix()

	_, err := h.rabbitMQActions.Ping(conn)

	if err != nil {
		c.AbortWithStatus(http.StatusBadGateway)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"uptime": uptime,
		})
	}
}
