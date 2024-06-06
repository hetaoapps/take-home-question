package main

import (
	"context"
	"database/sql"
	"log"
	"main/db"
	"main/handlers"
	"main/services"
	"main/utils"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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
			db.NewPostgresDB,
			utils.NewRedisClient,
			utils.NewHTTPClient,
			utils.NewOpenAIClient,
			services.NewUberEatsService,
			handlers.NewRecommendationHandler,
		),
		fx.Invoke(bootstrap),
	)

	app.Run()
}

func bootstrap(router *gin.Engine, handler *handlers.RecommendationHandler, db *sql.DB, redisClient *redis.Client) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := db.Ping()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	} else {
		log.Println("Connected to the database successfully")
	}

	err = utils.Ping(redisClient, ctx)
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	} else {
		log.Println("Connected to redis successfully")
	}

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		cancel()
	}()

	router.GET("/recommendations", func(c *gin.Context) {
		handler.GetRecommendations(c)
	})
	router.Run(":8080")
}
