package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/taheri24/helitask/pkg/config"
	"github.com/taheri24/helitask/pkg/di"
	"github.com/taheri24/helitask/pkg/domain"
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

	db, err := di.ProvideDB(cfg)
	if err != nil {
		slog.Error("failed to connect to database", slog.Any("err", err))
		os.Exit(1)
	}

	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("failed to access underlying sql.DB", slog.Any("err", err))
		os.Exit(1)
	}
	defer sqlDB.Close()

	if err := db.AutoMigrate(&domain.TodoItem{}); err != nil {
		slog.Error("failed to run migrations", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Info("database migrations applied successfully", slog.String("env", env))
}
