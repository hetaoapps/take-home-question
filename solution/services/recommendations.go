package services

import (
	"context"
	"fmt"
	"main/models"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func (s *UberEatsService) GetRecommendations(prompt string) ([]models.Recommendation, error) {
	food, err := s.GetFoodItemFromPrompt(prompt)
	var recommendations []models.Recommendation
	if err != nil {
		return recommendations, fmt.Errorf("failed to get food item from prompt: %w", err)
	}

	// check if the food recommendation is in the cache
	// go over the cache keys

	// check if the food recommendation is in the database

	// get the food recommendation from Uber Eats
	recommendations, err = s.GetRecommendationsFromUberEats(food)
	if err != nil {
		return recommendations, fmt.Errorf("failed to get recommendations from Uber Eats: %w", err)
	}

	// add the food recommendation to the cache

	// add the food recommendation to the database

	return recommendations, nil
}

func (s *UberEatsService) GetFoodItemFromPrompt(prompt string) (string, error) {
	promptContent, err := os.ReadFile("./prompts/foodItem.txt")
	if err != nil {
		return "", fmt.Errorf("failed to read prompt file: %w", err)
	}

	promptString := string(promptContent)
	promptToLLM := strings.Replace(promptString, "{{{userPrompt}}}", prompt, -1)

	resp, err := s.openaiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: promptToLLM,
				},
			},
		},
	)
	if err != nil {
		return "", fmt.Errorf("failed to get chat completion: %w", err)
	}

	return resp.Choices[0].Message.Content, nil
}
