package service

import (
	"engine/internal/storage/postgresql"
	"engine/internal/storage/redis"
)

type UserService struct {
}

func NewUserService(db *postgresql.DBStorage, rdb *redis.RedisStorage) *UserService {
	return &UserService{}
}
