package services

import (
	"database/sql"
	"net/http"

	"github.com/go-redis/redis/v8"
	openai "github.com/sashabaranov/go-openai"
)

type UberEatsService struct {
	openaiClient *openai.Client
	httpClient   *http.Client
	redis        *redis.Client
	db           *sql.DB
}

func NewUberEatsService(openaiClient *openai.Client, httpClient *http.Client, redis *redis.Client, db *sql.DB) *UberEatsService {
	return &UberEatsService{
		openaiClient: openaiClient,
		httpClient:   httpClient,
		redis:        redis,
		db:           db,
	}
}
