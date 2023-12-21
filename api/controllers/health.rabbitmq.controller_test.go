package controllers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"ppm3/go-aim-rest-api/configs"
	"ppm3/go-aim-rest-api/pkg/rabbitmq"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/mock"
)

func mockRabbitMQSetUp(mc *amqp.Connection, mP bool, errCon error, errPing error) rabbitmq.RabbitMQActionsI {
	var m *rabbitmq.MockRabbitMQActions = new(rabbitmq.MockRabbitMQActions)

	m.On("Connect", mock.Anything).Return(mc, errCon)
	m.On("Ping", mock.Anything).Return(mP, errPing)

	return m
}

func TestHealthRabbitMQController_CheckHealthRabbitMQ(t *testing.T) {
	type fields struct {
		ctx             context.Context
		rabbitConnect   *amqp.Connection
		configParams    *configs.ServerConfig
		rabbitMQActions rabbitmq.RabbitMQActionsI
	}
	type args struct {
		c      *gin.Context
		status int
	}

	ctx := context.Background()
	mc := &amqp.Connection{}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Successful request with status ok",
			fields: fields{
				ctx:             ctx,
				rabbitConnect:   mc,
				configParams:    &configs.ServerConfig{},
				rabbitMQActions: mockRabbitMQSetUp(mc, true, nil, nil),
			},
			args: args{
				c:      &gin.Context{},
				status: http.StatusOK,
			},
		},
		{
			name: "Ping request with status bad gateway",
			fields: fields{
				ctx:             ctx,
				rabbitConnect:   mc,
				configParams:    &configs.ServerConfig{},
				rabbitMQActions: mockRabbitMQSetUp(mc, false, nil, errors.New("mock error")),
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

			h := &HealthRabbitMQController{
				ctx:             tt.fields.ctx,
				rabbitConnect:   tt.fields.rabbitConnect,
				configParams:    tt.fields.configParams,
				rabbitMQActions: tt.fields.rabbitMQActions,
			}
			router.GET("/test", h.CheckHealthRabbitMQ)

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

func TestNewHealthRabbitMQController(t *testing.T) {
	type args struct {
		ctx context.Context
		r   *amqp.Connection
		rMQ rabbitmq.RabbitMQActionsI
		p   configs.ServerConfig
	}

	ctx := context.Background()
	rMQ := mockRabbitMQSetUp(&amqp.Connection{}, true, nil, nil)

	tests := []struct {
		name string
		args args
		want *HealthRabbitMQController
	}{
		{
			name: "Successful creation of HealthRabbitMQController",
			args: args{
				ctx: ctx,
				r:   &amqp.Connection{},
				p:   configs.ServerConfig{},
				rMQ: rMQ,
			},
			want: &HealthRabbitMQController{
				ctx:             ctx,
				rabbitConnect:   &amqp.Connection{},
				configParams:    &configs.ServerConfig{},
				rabbitMQActions: rMQ,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHealthRabbitMQController(tt.args.ctx, tt.args.r, tt.args.rMQ, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHealthRabbitMQController() = %v, want %v", got, tt.want)
			}
		})
	}
}
