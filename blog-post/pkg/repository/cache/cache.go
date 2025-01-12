package cache

import "github.com/redis/go-redis/v9"

type RedisConfig struct {
	Address string
}

func NewRedisCache(cfg RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: "",
		DB:       0,
	})

	return rdb
}
