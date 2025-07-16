package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port     string
	Database DatabaseConfig
	App      AppConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type AppConfig struct {
	Name    string
	Version string
	Env     string
}

func LoadConfig() *Config {
	return &Config{
		Port: getEnv("PORT", "3000"),
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			Name:     getEnv("DB_NAME", "telegramity"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
		},
		App: AppConfig{
			Name:    getEnv("APP_NAME", "Telegramity"),
			Version: getEnv("APP_VERSION", "1.0.0"),
			Env:     getEnv("APP_ENV", "development"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
