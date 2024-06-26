package redismq

import (
	"github.com/atauov/image-converter/internal/config"
	"github.com/go-redis/redis/v8"
)

type RedisMQ struct {
	Client *redis.Client
}

func NewRedisMQ(client *redis.Client) *RedisMQ {
	return &RedisMQ{Client: client}
}

func NewRedisClient(cfg *config.RedisServer) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Address,
	})

	return client
}
