package server

import (
	"context"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/keshu12345/truecaller/config"
	"go.uber.org/fx"
)

func TestInitializeServer(t *testing.T) {
	cfg := &config.Configuration{
		Server: config.Server{
			RestServicePort: 8080,
			ReadTimeout:     10,
			WriteTimeout:    10,
			IdleTimeout:     10,
		},
		EnvironmentName: "test",
	}

	// Create a test router
	router := gin.Default()

	// Create an fx application that includes the InitializeServer function
	app := fx.New(
		fx.Provide(func() *config.Configuration { return cfg }),
		fx.Provide(func() *gin.Engine { return router }),
		fx.Invoke(InitializeServer),
	)

	// Run the application
	go func() {
		if err := app.Start(context.Background()); err != nil {
			t.Fatalf("Failed to start application: %v", err)
		}
	}()

	// Sleep briefly to allow the goroutine to start
	time.Sleep(1 * time.Second)

	// Stop the application
	if err := app.Stop(context.Background()); err != nil {
		t.Fatalf("Failed to stop application: %v", err)
	}
}
