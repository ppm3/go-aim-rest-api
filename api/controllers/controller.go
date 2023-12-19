package controllers

import (
	"context"
	"ppm3/go-aim-rest-api/configs"
	"ppm3/go-aim-rest-api/pkg/databases"
)

type Controllers struct {
	PingController          *PingController
	HealthController        *HealthController
	HealthMongoDBController *HealthMongoDBController
}

func NewControllers(ctx context.Context, c *databases.Clients, m databases.MongoDBActionsI, p configs.ServerConfig) *Controllers {
	return &Controllers{
		PingController:          NewPingController(ctx, p),
		HealthController:        NewHealthController(ctx, p),
		HealthMongoDBController: NewHealthMongoDBController(ctx, c, m, p),
	}
}
