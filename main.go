package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/taheri24/helitask/pkg/config"
	"github.com/taheri24/helitask/pkg/di"
	"github.com/taheri24/helitask/pkg/logger"
	"github.com/taheri24/helitask/pkg/ports/storage"
	"github.com/taheri24/helitask/pkg/server"
	"go.uber.org/fx"
)

// main initializes and starts the application using dependency injection with Fx
func main() {
	// Load configuration based on the environment
	env := os.Getenv("APP_ENV") // Fetch environment (default to development if not set)
	if env == "" {
		env = "development"
	}

	// Load config for the specific environment
	cfg, err := config.LoadConfig(env)
	if err != nil {
		slog.Error("Error loading configuration", slog.Any("err", err))
		os.Exit(1)
	}

	defaultLogger := logger.New("app.log")
	appRoot := gin.Default()
	db, err := di.ProvideDB(cfg)
	if err != nil {
		slog.Error("di.ProvideDB failed", slog.Any("err", err))
		return
	}
	app := fx.New(
		fx.Supply(cfg, db, appRoot, defaultLogger),
		storage.Module,
		fx.Invoke(server.StartServer),
	)

	if err := app.Start(context.Background()); err != nil {
		slog.Error("Failed to start application", slog.Any("err", err))
		os.Exit(1)
	}

	defer app.Stop(context.Background())
}
