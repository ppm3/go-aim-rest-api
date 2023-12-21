package controllers

import (
	"context"
	"ppm3/go-aim-rest-api/configs"
	"ppm3/go-aim-rest-api/pkg/databases"
	"ppm3/go-aim-rest-api/pkg/rabbitmq"
	redisAction "ppm3/go-aim-rest-api/pkg/redis"

	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
)

type Controllers struct {
	PingController           *PingController
	HealthController         *HealthController
	HealthMySQLController    *HealthMySQLController
	HealthMongoDBController  *HealthMongoDBController
	HealthRabbitMQController *HealthRabbitMQController
	HealthRedisController    *HealthRedisController
}

func NewControllers(
	ctx context.Context,
	c *databases.Clients,
	rMQ *amqp.Connection,
	r *redis.Client,
	mda databases.MongoDBActionsI,
	msa databases.MySQLActionsI,
	rma rabbitmq.RabbitMQActionsI,
	ra redisAction.RedisActionsI,
	p configs.ServerConfig,
) *Controllers {
	return &Controllers{
		PingController:           NewPingController(ctx, p),
		HealthController:         NewHealthController(ctx, p),
		HealthMySQLController:    NewHealthMySQLController(ctx, c, msa, p),
		HealthMongoDBController:  NewHealthMongoDBController(ctx, c, mda, p),
		HealthRabbitMQController: NewHealthRabbitMQController(ctx, rMQ, rma, p),
		HealthRedisController:    NewHealthRedisController(ctx, r, ra, p),
	}
}
