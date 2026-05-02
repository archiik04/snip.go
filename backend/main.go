package main

import (
	"log"
	"os"

	"snip-go/handlers"
	"snip-go/store"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	redisURL := getEnv("REDIS_URL", "redis://localhost:6379")
	baseURL := getEnv("BASE_URL", "http://localhost:8080")
	port := getEnv("PORT", "8080")

	redisStore := store.NewRedisStore(redisURL)
	urlHandler := handlers.NewURLHandler(redisStore, baseURL)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
	}))

	api := r.Group("/api")
	{
		api.POST("/shorten", urlHandler.Shorten)
		api.GET("/stats/:code", urlHandler.Stats)
	}

	r.GET("/:code", urlHandler.Redirect)

	log.Printf("Server running on :%s", port)
	r.Run(":" + port)
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
