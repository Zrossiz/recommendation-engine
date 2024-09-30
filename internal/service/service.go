package service

import (
	"engine/internal/storage/postgresql"
	"engine/internal/storage/redis"
)

type Service struct {
	Category *CategoryService
	User     *UserService
	Content  *ContentService
}

func New(db *postgresql.DBStorage, rdb *redis.RedisStorage) *Service {
	return &Service{
		Category: NewCategoryService(db, rdb),
		User:     NewUserService(db, rdb),
		Content:  NewContentService(db, rdb),
	}
}
