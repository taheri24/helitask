package di

import (
	"regexp"
	"strconv"

	"log/slog"

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
func ProvideDB(cfg *config.Config) (*gorm.DB, error) {
	log := logger.New("")
	log.Verbose("Connecting to database with DSN:", cfg.DB.DSN)

	db, err := postgres.NewDB(cfg.DB.DSN)
	if err != nil {
		log.Error("Failed to connect to database", err)
		return nil, err
	}
	sqlDb, _ := db.DB()
	if err := sqlDb.Ping(); err != nil {
		log.Error("Database connection failed", err)
		return nil, err
	}

	var version string
	if err := db.Raw("SELECT version();").Scan(&version).Error; err != nil {
		log.Error("Failed to fetch PostgreSQL version", err)
		return nil, err
	}
	log.Verbose("Connected to PostgreSQL", "version", version)

	if isVersionGreaterThan15(version) {
		slog.Warn("Warning: PostgreSQL major version is greater than 15, which may cause compatibility issues.", "version", version)
	}

	return db, nil
}

// isVersionGreaterThan15 checks if the PostgreSQL version is greater than 15
func isVersionGreaterThan15(version string) bool {
	re := regexp.MustCompile(`PostgreSQL (\d+)\.\d+\.\d+`)
	matches := re.FindStringSubmatch(version)

	if len(matches) > 1 {
		majorVersion, err := strconv.Atoi(matches[1])
		if err != nil {
			slog.Error("Failed to parse PostgreSQL version", slog.Any("err", err))
			return false
		}
		return majorVersion > 15
	}
	slog.Warn("Unable to parse PostgreSQL version string.")
	return false
}
