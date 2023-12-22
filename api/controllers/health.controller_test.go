package controllers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"ppm3/go-aim-rest-api/configs"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCheckHealth(t *testing.T) {
	router := gin.Default()
	c := context.Background()

	cp := &configs.ServerConfig{}
	controller := NewHealthController(c, *cp)
	router.GET("/test", controller.Ping)

	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}

func TestNewHealthController(t *testing.T) {
	type args struct {
		ctx context.Context
		p   configs.ServerConfig
	}
	tests := []struct {
		name string
		args args
		want *HealthController
	}{
		{
			name: "Successful creation of HealthController",
			args: args{
				ctx: context.Background(),
				p:   configs.ServerConfig{},
			},
			want: &HealthController{
				ctx:          context.Background(),
				configParams: &configs.ServerConfig{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHealthController(tt.args.ctx, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHealthController() = %v, want %v", got, tt.want)
			}
		})
	}
}
