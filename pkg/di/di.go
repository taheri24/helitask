package di

import (
	"github.com/taheri24/helitask/pkg/config"
	"github.com/taheri24/helitask/pkg/logger"
	"github.com/taheri24/helitask/pkg/ports/storage/postgres"
	"gorm.io/gorm"
)

// ProvideConfig loads configuration settings from the environment
func ProvideConfig() (*config.Config, error) {
	log := logger.New("")
	log.Verbose("Loading configuration settings")
	return config.LoadConfig("development")
}

// ProvideLogger provides a logger instance
func ProvideLogger(cfg *config.Config) logger.Logger {
	log := logger.New("")
	log.Verbose("Providing logger instance")
	return log
}

// ProvideDB establishes the database connection and provides a TodoRepository
func ProvideDB(cfg *config.Config, logger logger.Logger) (*gorm.DB, error) {
	if cfg == nil {
		panic("cfg==nil")
	}
	db, err := postgres.NewDB(cfg.DB.DSN)
	if err != nil {
		logger.Error("Failed to connect to database", err)
		return nil, err
	}
	sqlDb, _ := db.DB()
	if err := sqlDb.Ping(); err != nil {
		logger.Error("Database connection failed", err)
		return nil, err
	}

	return db, nil
}
