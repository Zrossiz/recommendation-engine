package app

import (
	"engine/internal/config"
	"engine/internal/service"
	"engine/internal/storage/postgresql"
	"engine/internal/storage/redis"
	"engine/internal/transport/http/handler"
	"engine/internal/transport/http/router"
	"engine/pkg/logger"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func Start() {
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
	db := postgresql.New(dbClient, log)
	log.Info("db ready")

	redisClient := redis.Connect(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
	_ = redis.New(redisClient)
	log.Info("redis ready")

	serv := service.New(service.Storage{
		UserStore:             db.UserStore,
		UserInteractionsStore: db.UserInteractionsStore,
		ContentStore:          db.ContentStore,
		CategoryStore:         db.CategoryStore,
	}, log, cfg)
	h := handler.New(serv)
	r := router.New(h)

	srv := &http.Server{
		Addr:    cfg.RunAddr,
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Failed to start server", zap.Error(err))
	}
}
