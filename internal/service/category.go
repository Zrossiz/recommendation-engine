package service

import (
	"engine/internal/storage/postgresql"
	"engine/internal/storage/redis"
)

type CategoryService struct {
}

func NewCategoryService(db *postgresql.DBStorage, rdb *redis.RedisStorage) *CategoryService {
	return &CategoryService{}
}
