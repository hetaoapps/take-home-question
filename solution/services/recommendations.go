package services

import (
	"context"
	"encoding/json"
	"fmt"
	"main/models"
	"main/utils"
	"os"
	"strconv"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func (s *UberEatsService) GetRecommendations(ctx context.Context, prompt string) ([]models.Recommendation, error) {
	food, err := s.GetFoodItemFromPrompt(prompt)
	var recommendations []models.Recommendation
	if err != nil {
		return recommendations, fmt.Errorf("failed to get food item from prompt: %w", err)
	}

	// check if the food recommendation is in the cache
	// go over the cache keys
	fmt.Println("Getting recommendations from cache")
	hash, err := s.redis.HGetAll(ctx, "recommendations").Result()
	if err != nil {
		return recommendations, fmt.Errorf("failed to get recommendations from cache: %w", err)
	}
	// Iterate over each key-value pair in the hash
	thresholdStr := os.Getenv("FOOD_ITEM_SIMILARITY_THRESHOLD")
	threshold, err := strconv.ParseFloat(thresholdStr, 64)
	if err != nil {
		return recommendations, fmt.Errorf("failed to parse threshold: %w", err)
	}

	for key, value := range hash {
		similarity := utils.GetSimilarity(food, key)
		if similarity >= threshold {
			// unmarshal the value into a recommendation struct
			err = json.Unmarshal([]byte(value), &recommendations)
			if err != nil {
				return recommendations, fmt.Errorf("failed to unmarshal recommendation: %w", err)
			}
			return recommendations, nil
		}
	}

	// check if the food recommendation is in the database
	fmt.Println("Getting recommendations from database")
	rows, err := s.db.QueryContext(ctx, "SELECT * FROM recommendations WHERE food = $1", food)
	if err != nil {
		return recommendations, fmt.Errorf("failed to query database: %w", err)
	}
	defer rows.Close()
	// go over the rwos and check if the food recommendation is in the database
	for rows.Next() {
		var current_food string
		if err := rows.Scan(&current_food); err != nil {
			return recommendations, fmt.Errorf("failed to scan row: %w", err)
		}
		similarity := utils.GetSimilarity(food, current_food)
		if similarity >= threshold {
			s.constructRecommendationsFromDB(storeID)
		}
	}

	// get the food recommendation from Uber Eats
	recommendations, err = s.GetRecommendationsFromUberEats(food)
	if err != nil {
		return recommendations, fmt.Errorf("failed to get recommendations from Uber Eats: %w", err)
	}

	// add the food recommendation to the cache
	recommendationsJSON, err := json.Marshal(recommendations)
	if err != nil {
		return recommendations, fmt.Errorf("failed to marshal recommendations: %w", err)
	}

	fmt.Println(string(recommendationsJSON))
	err = s.redis.HSet(ctx, "recommendations", food, recommendationsJSON).Err()
	if err != nil {
		return recommendations, fmt.Errorf("failed to add recommendation to cache: %w", err)
	}

	// add the food recommendation to the database
	s.AddRecommendationsToDB(recommendations, food)

	return recommendations, nil
}

func (s *UberEatsService) constructRecommendationsFromDB(recommendationID string, storeID string) error {
	// get store from the database
	store := models.RecommendStore{}
	err := s.db.QueryRow("SELECT * FROM stores WHERE id = $1", storeID).Scan(&store.ID, &store.Name)
	if err != nil {
		return fmt.Errorf("failed to get store from database: %w", err)
	}

	// get menu items from the database
	rows, err := s.db.Query("SELECT * FROM menu_items WHERE store_id = $1", storeID)
	if err != nil {
		return fmt.Errorf("failed to get menu items from database: %w", err)
	}
	defer rows.Close()

	var menuItems []models.MenuItem
	for rows.Next() {
		menuItem := models.MenuItem{}
		if err := rows.Scan(&menuItem.ID, &menuItem.Name, &menuItem.Price); err != nil {
			return fmt.Errorf("failed to scan row: %w", err)
		}
		menuItems = append(menuItems, menuItem)
	}

	store.MenuItem = menuItems

	// add the recommendation to the list of recommendations
	recommendation := models.Recommendation{
		RecommendationID: recommendationID,
		Store:            store,
	}
	recommendations = append(recommendations, recommendation)
}

func (s *UberEatsService) AddRecommendationsToDB(recommendations []models.Recommendation, food string) error {
	// insert store
	for _, recommendation := range recommendations {
		// insert recommendation
		_, err := s.db.Exec("INSERT INTO Recommendations (query, store_id) VALUES ($1, $2)", food, recommendation.RecommendationID)
		if err != nil {
			return fmt.Errorf("failed to insert recommendation: %w", err)
		}
		// insert store
		_, err = s.db.Exec("INSERT INTO Stores (id, name) VALUES ($1, $2)", recommendation.RecommendationID, recommendation.Store.Name)
		if err != nil {
			return fmt.Errorf("failed to insert store: %w", err)
		}
		// insert menu items
		for _, menuItem := range recommendation.Store.MenuItem {
			_, err = s.db.Exec("INSERT INTO MenuItems (id, store_id, name, price) VALUES ($1, $2, $3, $4)", menuItem.ID, recommendation.Store.ID, recommendation.Store.Name, menuItem.Price)
			if err != nil {
				return fmt.Errorf("failed to insert menu item: %w", err)
			}
		}
	}
	return nil
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
