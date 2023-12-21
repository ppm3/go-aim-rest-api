package controllers

import (
	"context"
	"ppm3/go-aim-rest-api/configs"
	"ppm3/go-aim-rest-api/pkg/databases"
	"ppm3/go-aim-rest-api/pkg/rabbitmq"

	"github.com/streadway/amqp"
)

type Controllers struct {
	PingController           *PingController
	HealthController         *HealthController
	HealthMySQLController    *HealthMySQLController
	HealthMongoDBController  *HealthMongoDBController
	HealthRabbitMQController *HealthRabbitMQController
}

func NewControllers(
	ctx context.Context,
	c *databases.Clients,
	r *amqp.Connection,
	mda databases.MongoDBActionsI,
	msa databases.MySQLActionsI,
	ra rabbitmq.RabbitMQActionsI,
	p configs.ServerConfig,
) *Controllers {
	return &Controllers{
		PingController:           NewPingController(ctx, p),
		HealthController:         NewHealthController(ctx, p),
		HealthMySQLController:    NewHealthMySQLController(ctx, c, msa, p),
		HealthMongoDBController:  NewHealthMongoDBController(ctx, c, mda, p),
		HealthRabbitMQController: NewHealthRabbitMQController(ctx, r, ra, p),
	}
}
