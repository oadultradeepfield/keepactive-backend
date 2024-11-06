package config

import (
	"os"
	"strings"
)

type Config struct {
    Port            string
    DatabaseURL     string
    JWTSecret       string
    AllowedOrigins  []string
    Environment     string
}

func LoadConfig() *Config {
    return &Config{
        Port:           getEnvOrDefault("PORT", "8080"),
        DatabaseURL:    os.Getenv("DATABASE_URL"),
        JWTSecret:     os.Getenv("JWT_SECRET"),
        AllowedOrigins: strings.Split(os.Getenv("ALLOWED_ORIGINS"), ","),
        Environment:    getEnvOrDefault("GO_ENV", "development"),
    }
}

func getEnvOrDefault(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}