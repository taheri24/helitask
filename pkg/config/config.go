package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Config represents the application configuration
type Config struct {
	DB     DatabaseConfig
	Server ServerConfig
}

// DatabaseConfig holds database-related settings
type DatabaseConfig struct {
	DSN string
}

// ServerConfig holds the server-related settings
type ServerConfig struct {
	Port string
}

// LoadConfig loads the configuration from environment files
// It attempts to load an environment-specific `.env` file, with a fallback to `.env` as default.
func LoadConfig(env string) (*Config, error) {
	// Load the environment-specific file (e.g., .env.production, .env.development)
	envFile := fmt.Sprintf(".env.%s", env)
	if err := godotenv.Load(envFile); err != nil {
		// If specific env file not found, fall back to default `.env` file
		if err := godotenv.Load(".env"); err != nil {
			return nil, fmt.Errorf("Error loading .env file")
		}
	}

	// Use viper to load environment variables
	viper.SetConfigFile(fmt.Sprintf(".env.%s", env))
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Error loading .env file: %s", err)
	}

	// Set default values if necessary
	viper.SetDefault("DB_DSN", "localhost:5432")
	viper.SetDefault("PORT", "8080")

	return &Config{
		DB: DatabaseConfig{
			DSN: viper.GetString("DB_DSN"),
		},
		Server: ServerConfig{
			Port: viper.GetString("PORT"),
		},
	}, nil
}
