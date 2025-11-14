package config

import (
	"fmt"
	"os"

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

func fileSize(fn string) int64 {
	st, err := os.Stat(fn)
	if err != nil {
		return 0
	}

	return st.Size()
}

// LoadConfig loads the configuration from environment files
// It attempts to load an environment-specific `.env` file, with a fallback to `.env` as default.
func LoadConfig(env string) (*Config, error) {
	// Load the environment-specific file (e.g., .env.production, .env.development)
	envFile := fmt.Sprintf(".env.%s", env)
	if fileSize(envFile) > 0 {
		viper.SetConfigFile(envFile)

	}
	envFile = ".env"
	if fileSize(envFile) > 0 {
		viper.SetConfigFile(envFile)
	}

	// Use viper to load environment variables
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %s", err)
	}
	viper.AutomaticEnv()
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
