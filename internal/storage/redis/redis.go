package redis

import "github.com/redis/go-redis/v9"

type RedisStorage struct {
	db *redis.Client
}

func Connect(addr string, password string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return rdb
}

func New(rdb *redis.Client) *RedisStorage {
	return &RedisStorage{db: rdb}
}
