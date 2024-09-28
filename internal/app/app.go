package app

import (
	"engine/internal/config"
	"engine/internal/storage/postgresql"
	"engine/internal/storage/redis"
	"engine/pkg/logger"
	"fmt"

	"go.uber.org/zap"
)

func Start() {
	fmt.Println("start")
	cfg, err := config.Init()
	if err != nil {
		fmt.Printf("parse config error: %v\n", err.Error())
	}
	zap.S().Info("config ready")

	zapLogger, err := logger.New(cfg.LogLevel)
	if err != nil {
		zap.S().Fatalf("init logger error: %v", err)
	}
	log := zapLogger.ZapLogger
	log.Info("logger ready")

	dbClient, err := postgresql.Connect(cfg.DBDSN)
	if err != nil {
		log.Fatal(err.Error())
	}
	db := postgresql.New(dbClient)
	log.Info("db ready")

	redisClient := redis.Connect(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
	rdb := redis.New(redisClient)
	log.Info("redis ready")

}
