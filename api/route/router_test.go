package route

import (
	"context"
	"net/http"
	"net/http/httptest"
	"ppm3/go-aim-rest-api/api/controllers"
	"ppm3/go-aim-rest-api/configs"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSetupRouter(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Initialize your controllers and pass them to the SetupRouter function
	var c controllers.Controllers = *controllers.NewControllers(
		context.Background(),
		configs.ServerConfig{},
	)

	// Call the SetupRouter function to set up the routes
	SetupRouter(context.Background(), c, router)

	// Define the paths you want to test
	paths := []string{
		"/",
		"/health-check",
		"/ping",
	}

	// Iterate over the paths and test the HTTP status code
	for _, path := range paths {
		req, _ := http.NewRequest("GET", path, nil)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d for path %s", http.StatusOK, resp.Code, path)
		}
	}
}
