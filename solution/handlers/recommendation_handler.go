package handlers

import (
	"net/http"

	"main/services"

	"github.com/gin-gonic/gin"
)

type RecommendationHandler struct {
	uberEatsService *services.UberEatsService
}

func NewRecommendationHandler(uberEatsService *services.UberEatsService) *RecommendationHandler {
	return &RecommendationHandler{uberEatsService: uberEatsService}
}

func (h *RecommendationHandler) GetRecommendations(c *gin.Context) {
	prompt := c.Query("prompt")
	if prompt == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "prompt query parameter is required"})
		return
	}

	food, err := h.uberEatsService.GetFoodItemFromPrompt(prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process prompt"})
		return
	}

	recommendations, err := h.uberEatsService.GetRecommendationsFromUberEats(food)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get recommendations"})
		return
	}

	c.JSON(http.StatusOK, recommendations)
}
