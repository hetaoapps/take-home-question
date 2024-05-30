package main

import (
	"log"
	"main/handlers"

	"main/services"
	"main/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fx.New(
		fx.Provide(
			gin.New,
			utils.NewHTTPClient,
			utils.NewOpenAIClient,
			services.NewUberEatsService,
			handlers.NewRecommendationHandler,
		),
		fx.Invoke(bootstrap),
	)

	app.Run()
}

func bootstrap(router *gin.Engine, handler *handlers.RecommendationHandler) {
	router.GET("/recommendations", handler.GetRecommendations)
	router.Run("localhost:8080")
}
