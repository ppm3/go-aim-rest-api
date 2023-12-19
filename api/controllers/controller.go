package controllers

import (
	"context"
	"ppm3/go-aim-rest-api/configs"
)

func NewControllers(ctx context.Context, p configs.ServerConfig) *Controllers {
	return &Controllers{
		PingController:   NewPingController(ctx, p),
		HealthController: NewHealthController(ctx, p),
	}
}
