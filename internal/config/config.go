package config

import (
	"os"
	"strconv"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Qdrant   QdrantConfig
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

// QdrantConfig holds Qdrant vector database configuration
type QdrantConfig struct {
	Host string
	Port string
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "admin"),
			Password: getEnv("DB_PASSWORD", "admin123"),
			Name:     getEnv("DB_NAME", "retailBrain"),
		},
		Qdrant: QdrantConfig{
			Host: getEnv("QDRANT_HOST", "localhost"),
			Port: getEnv("QDRANT_PORT", "6333"),
		},
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}

// getEnvAsInt gets an environment variable as integer or returns a default value
func getEnvAsInt(key string, fallback int) int {
	if val, ok := os.LookupEnv(key); ok {
		if intVal, err := strconv.Atoi(val); err == nil {
			return intVal
		}
	}
	return fallback
}
