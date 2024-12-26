package redis

import (
	"github.com/amankumarsingh77/cmr/config"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func NewRedisClient(config *config.Config) *redis.Client {
	redisHost := config.Redis.RedisAddr

	log.Println("redisHost", config.Redis)

	if redisHost == "" {
		redisHost = ":6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:         redisHost,
		MinIdleConns: config.Redis.MinIdleConns,
		PoolSize:     config.Redis.PoolSize,
		PoolTimeout:  time.Duration(config.Redis.PoolTimeout) * time.Second,
		Password:     config.Redis.RedisPassword,
		DB:           config.Redis.DB,
	})
	return client
}
