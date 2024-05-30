package cache

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	})
}

func Ping(client *redis.Client, ctx context.Context) error {
	_, err := client.Ping(ctx).Result()
	return err
}
