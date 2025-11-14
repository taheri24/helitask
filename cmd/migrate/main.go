package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/taheri24/helitask/pkg/config"
	"github.com/taheri24/helitask/pkg/di"
	"github.com/taheri24/helitask/pkg/domain"
	"github.com/taheri24/helitask/pkg/logger"
)

func main() {
	envFlag := flag.String("env", "", "Environment to load configuration from")
	flag.Parse()

	env := *envFlag
	if env == "" {
		env = os.Getenv("APP_ENV")
	}
	if env == "" {
		env = "development"
	}

	cfg, err := config.LoadConfig(env)
	if err != nil {
		slog.Error("failed to load configuration", slog.Any("err", err))
		os.Exit(1)
	}
	_logger := logger.NewSlogger(slog.Default())
	db, err := di.ProvideDB(cfg, _logger)
	if err != nil {
		slog.Error("failed to connect to database", slog.Any("err", err))
		os.Exit(1)
	}

	if err := db.AutoMigrate(&domain.TodoItem{}); err != nil {
		slog.Error("failed to run migrations", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Info("database migrations applied successfully", slog.String("env", env))
}
