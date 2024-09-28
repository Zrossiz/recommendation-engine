package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RunAddr          string
	DBDSN            string
	JWTAccessSecret  string
	JWTRefreshSecret string
	LogLevel         string
	RedisAddr        string
	RedisPassword    string
	RedisDB          int
}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%s not provided", key)
	}
	return value, nil
}

func getEnvAsInt(key string) (int, error) {
	valueStr, err := getEnv(key)
	if err != nil {
		return 0, err
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, fmt.Errorf("%s must be a valid integer", key)
	}
	return value, nil
}

func Init() (*Config, error) {
	_ = godotenv.Load(".env")

	cfg := Config{}

	var err error

	if cfg.DBDSN, err = getEnv("DB_DSN"); err != nil {
		return nil, err
	}
	if cfg.RedisAddr, err = getEnv("REDIS_ADDRESS"); err != nil {
		return nil, err
	}
	if cfg.RedisPassword, err = getEnv("REDIS_PASSWORD"); err != nil {
		return nil, err
	}
	if cfg.RedisDB, err = getEnvAsInt("REDIS_DB"); err != nil {
		return nil, err
	}
	if cfg.JWTAccessSecret, err = getEnv("JWT_ACCESS_SECRET"); err != nil {
		return nil, err
	}
	if cfg.JWTRefreshSecret, err = getEnv("JWT_REFRESH_SECRET"); err != nil {
		return nil, err
	}
	if cfg.LogLevel, err = getEnv("LOG_LEVEL"); err != nil {
		return nil, err
	}

	return &cfg, nil
}
