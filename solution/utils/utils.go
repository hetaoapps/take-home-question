package utils

import (
	"context"
	"net/http"
	"os"

	openai "github.com/sashabaranov/go-openai"

	"github.com/go-redis/redis/v8"
)

func NewHTTPClient() *http.Client {
	return &http.Client{}
}

func NewOpenAIClient() *openai.Client {
	return openai.NewClient(os.Getenv("OPENAI_API_KEY"))
}

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
	})
}

func Ping(client *redis.Client, ctx context.Context) error {
	_, err := client.Ping(ctx).Result()
	return err
}
