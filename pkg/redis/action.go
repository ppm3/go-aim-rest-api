package redisAction

import (
	"context"
	"fmt"
	"ppm3/go-aim-rest-api/configs"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type RedisActionsI interface {
	Connect() (*redis.Client, error)
	Ping(rc *redis.Client) (bool, error)
}

type RedisAction struct {
	ctx    context.Context
	params *configs.RedisConfig
}

func NewRedisConnect(ctx context.Context, params *configs.RedisConfig) *RedisAction {
	return &RedisAction{
		ctx:    ctx,
		params: params,
	}
}

func (r *RedisAction) Connect() (*redis.Client, error) {
	dbStr := r.params.Database
	db, _ := strconv.Atoi(dbStr)

	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", r.params.Host, r.params.Port),
		Password:     r.params.Password,
		DB:           db,
		DialTimeout:  time.Duration(r.params.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(r.params.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(r.params.WriteTimeout) * time.Second,
		PoolSize:     r.params.PoolSize,
		MinIdleConns: r.params.MinIdleConns,
		IdleTimeout:  time.Duration(r.params.IdleTimeout) * time.Minute,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (r *RedisAction) Ping(rc *redis.Client) (bool, error) {

	pong, err := rc.Ping().Result()
	if err != nil {
		return false, err
	}

	return pong == "PONG", nil
}
