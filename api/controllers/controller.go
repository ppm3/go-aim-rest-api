package controllers

import (
	"context"
	"ppm3/go-aim-rest-api/configs"
	"ppm3/go-aim-rest-api/pkg/databases"
)

type Controllers struct {
	PingController          *PingController
	HealthController        *HealthController
	HealthMySQLController   *HealthMySQLController
	HealthMongoDBController *HealthMongoDBController
}

func NewControllers(ctx context.Context, c *databases.Clients, mda databases.MongoDBActionsI, msa databases.MySQLActionsI, p configs.ServerConfig) *Controllers {
	return &Controllers{
		PingController:          NewPingController(ctx, p),
		HealthController:        NewHealthController(ctx, p),
		HealthMySQLController:   NewHealthMySQLController(ctx, c, msa, p),
		HealthMongoDBController: NewHealthMongoDBController(ctx, c, mda, p),
	}
}
