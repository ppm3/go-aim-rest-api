package controllers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"ppm3/go-aim-rest-api/configs"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPingController_Pong(t *testing.T) {
	router := gin.Default()
	c := context.Background()

	cp := &configs.ServerConfig{}
	controller := NewPingController(c, *cp)
	router.GET("/test", controller.Pong)

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
