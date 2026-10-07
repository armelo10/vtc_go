package config

import "os"

type Config struct {
	HTTPAddr    string
	DatabaseURL string
	RedisURL    string
	Environment string
}

func Load() Config {
	return Config{
		HTTPAddr:    env("HTTP_ADDR", ":8080"),
		DatabaseURL: env("DATABASE_URL", "postgres://vtc:vtc@localhost:5432/vtc?sslmode=disable"),
		RedisURL:    env("REDIS_URL", "redis://localhost:6379/0"),
		Environment: env("APP_ENV", "development"),
	}
}

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
