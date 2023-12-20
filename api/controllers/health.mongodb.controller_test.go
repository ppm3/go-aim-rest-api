package controllers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"ppm3/go-aim-rest-api/configs"
	"ppm3/go-aim-rest-api/pkg/databases"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

func mockMongoSetUp(mc *mongo.Client, mP bool, mD bool, errCon error, errPing error, errDB error) databases.MongoDBActionsI {
	var m *databases.MockMongoDBActions = new(databases.MockMongoDBActions)

	m.On("MongoConnect", mock.Anything).Return(mc, errCon)
	m.On("MongoPing", mock.Anything).Return(mP, errPing)
	m.On("MongoPingDB", mock.Anything).Return(mD, errDB)

	return m
}

func TestHealthMongoDBController_CheckHealthDB(t *testing.T) {
	type fields struct {
		ctx          context.Context
		clients      *databases.Clients
		mongoActions databases.MongoDBActionsI
		configParams *configs.ServerConfig
	}
	type args struct {
		c      *gin.Context
		status int
	}

	ctx := context.Background()
	mc := &mongo.Client{}

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
				mongoActions: mockMongoSetUp(mc, true, true, nil, nil, nil),
				configParams: &configs.ServerConfig{},
			},
			args: args{
				c:      &gin.Context{},
				status: http.StatusOK,
			},
		},
		{
			name: "Error request with status mongo ping failure",
			fields: fields{
				ctx:          ctx,
				clients:      &databases.Clients{},
				mongoActions: mockMongoSetUp(mc, false, true, nil, errors.New("mock error"), nil),
				configParams: &configs.ServerConfig{},
			},
			args: args{
				c:      &gin.Context{},
				status: http.StatusInternalServerError,
			},
		},
		{
			name: "Error request with status db ping failure",
			fields: fields{
				ctx:          ctx,
				clients:      &databases.Clients{},
				mongoActions: mockMongoSetUp(mc, true, false, nil, nil, errors.New("mock error")),
				configParams: &configs.ServerConfig{},
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

			h := &HealthMongoDBController{
				ctx:          tt.fields.ctx,
				clients:      tt.fields.clients,
				mongoActions: tt.fields.mongoActions,
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
