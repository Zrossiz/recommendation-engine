package service

import (
	"engine/internal/storage/postgresql"
	"engine/internal/storage/redis"
)

type ContentService struct {
}

func NewContentService(db *postgresql.DBStorage, rdb *redis.RedisStorage) *ContentService {
	return &ContentService{}
}
