package models

import (
	"context"
	"log"

	"github.com/aditya-amlesh-jha/url-shortener/server/config"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func InitRedis(cfg *config.Config) *redis.Client {
	redisDb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	if _, err := redisDb.Ping(ctx).Result(); err != nil {
		log.Fatalf("Failed to ping Redis :: %v", err)
	}

	return redisDb
}
