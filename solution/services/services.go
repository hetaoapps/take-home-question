package services

import (
	"net/http"

	openai "github.com/sashabaranov/go-openai"
)

type UberEatsService struct {
	openaiClient *openai.Client
	httpClient   *http.Client
}

func NewUberEatsService(openaiClient *openai.Client, httpClient *http.Client) *UberEatsService {
	return &UberEatsService{
		openaiClient: openaiClient,
		httpClient:   httpClient,
	}
}
