package controllers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"ppm3/go-aim-rest-api/configs"
	redisAction "ppm3/go-aim-rest-api/pkg/redis"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/mock"
)

func mockRedis(r *redis.Client, pResp bool, errClient error, errPing error) redisAction.RedisActionsI {
	var m *redisAction.MockRedisActions = new(redisAction.MockRedisActions)

	m.On("Connect", mock.Anything).Return(&redis.Client{}, errClient)
	m.On("Ping", mock.Anything).Return(pResp, errPing)

	return m
}

func TestHealthRedisController_CheckHealthRedis(t *testing.T) {
	type fields struct {
		ctx          context.Context
		redisClient  *redis.Client
		configParams *configs.ServerConfig
		redisActions redisAction.RedisActionsI
	}
	type args struct {
		c      *gin.Context
		status int
	}

	ctx := context.Background()
	rc := &redis.Client{}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Successful request with status ok",
			fields: fields{
				ctx:          ctx,
				redisClient:  rc,
				configParams: &configs.ServerConfig{},
				redisActions: mockRedis(rc, true, nil, nil),
			},
			args: args{
				c:      &gin.Context{},
				status: http.StatusOK,
			},
		},
		{
			name: "Failed request with status bad gateway",
			fields: fields{
				ctx:          ctx,
				redisClient:  rc,
				configParams: &configs.ServerConfig{},
				redisActions: mockRedis(rc, false, nil, errors.New("mock error")),
			},
			args: args{
				c:      &gin.Context{},
				status: http.StatusBadGateway,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.Default()

			h := &HealthRedisController{
				ctx:          tt.fields.ctx,
				redisClient:  tt.fields.redisClient,
				configParams: tt.fields.configParams,
				redisActions: tt.fields.redisActions,
			}

			router.GET("/test", h.CheckHealthRedis)

			req, err := http.NewRequest("GET", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}

			recorder := httptest.NewRecorder()

			router.ServeHTTP(recorder, req)

			if recorder.Code != tt.args.status {
				t.Errorf("Expected status code %d, but got %d", tt.args.status, recorder.Code)
			}
		})
	}
}

func TestNewHealthRedisController(t *testing.T) {
	type args struct {
		ctx context.Context
		r   *redis.Client
		rA  redisAction.RedisActionsI
		p   configs.ServerConfig
	}

	ctx := context.Background()
	rc := &redis.Client{}
	rA := mockRedis(rc, true, nil, nil)

	tests := []struct {
		name string
		args args
		want *HealthRedisController
	}{
		{
			name: "NewHealthRedisController",
			args: args{
				ctx: ctx,
				r:   rc,
				rA:  rA,
				p:   configs.ServerConfig{},
			},
			want: &HealthRedisController{
				ctx:          ctx,
				redisClient:  rc,
				configParams: &configs.ServerConfig{},
				redisActions: rA,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHealthRedisController(tt.args.ctx, tt.args.r, tt.args.rA, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHealthRedisController() = %v, want %v", got, tt.want)
			}
		})
	}
}
