package utils

import (
	"net/http"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func NewHTTPClient() *http.Client {
	return &http.Client{}
}

func NewOpenAIClient() *openai.Client {
	return openai.NewClient(os.Getenv("OPENAI_API_KEY"))
}
