package controllers

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"ppm3/go-aim-rest-api/configs"
	"ppm3/go-aim-rest-api/pkg/databases"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func mockMySQLSetUp(mc *sql.DB, mP bool, errCon error, errPing error) databases.MySQLActionsI {
	var m *databases.MockMySQLActions = new(databases.MockMySQLActions)

	m.On("MySQLConnect", mock.Anything).Return(mc, errCon)
	m.On("MySQLPing", mock.Anything).Return(mP, errPing)

	return m
}

func TestHealthMySQLController_CheckHealthDB(t *testing.T) {
	type fields struct {
		ctx          context.Context
		clients      *databases.Clients
		mysqlActions databases.MySQLActionsI
		configParams *configs.ServerConfig
	}
	type args struct {
		c      *gin.Context
		status int
	}

	ctx := context.Background()
	mc := &sql.DB{}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Successful request with status ok",
			fields: fields{
				ctx:          ctx,
				clients:      &databases.Clients{},
				mysqlActions: mockMySQLSetUp(mc, true, nil, nil),
				configParams: &configs.ServerConfig{},
			},
			args: args{
				c:      &gin.Context{},
				status: http.StatusOK,
			},
		},
		{
			name: "Failed request with status internal server error",
			fields: fields{
				ctx:          ctx,
				clients:      &databases.Clients{},
				mysqlActions: mockMySQLSetUp(mc, false, nil, errors.New("mock error")),
				configParams: &configs.ServerConfig{},
			},
			args: args{
				c:      &gin.Context{},
				status: http.StatusInternalServerError,
			},
		},
	}
	for _, tt := range tests {
		router := gin.Default()

		t.Run(tt.name, func(t *testing.T) {
			h := &HealthMySQLController{
				ctx:          tt.fields.ctx,
				clients:      tt.fields.clients,
				mysqlActions: tt.fields.mysqlActions,
				configParams: tt.fields.configParams,
			}

			router.GET("/test", h.CheckHealthDB)

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

func TestNewHealthMySQLController(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *databases.Clients
		m   databases.MySQLActionsI
		p   configs.ServerConfig
	}
	tests := []struct {
		name string
		args args
		want *HealthMySQLController
	}{
		{
			name: "Successful creation of HealthMySQLController",
			args: args{
				ctx: context.Background(),
				c:   &databases.Clients{},
				m:   &databases.MockMySQLActions{},
				p:   configs.ServerConfig{},
			},
			want: &HealthMySQLController{
				ctx:          context.Background(),
				clients:      &databases.Clients{},
				mysqlActions: &databases.MockMySQLActions{},
				configParams: &configs.ServerConfig{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHealthMySQLController(tt.args.ctx, tt.args.c, tt.args.m, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHealthMySQLController() = %v, want %v", got, tt.want)
			}
		})
	}
}
