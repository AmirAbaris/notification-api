package main

import (
	"context"
	"log"

	"github.com/AmirAbaris/notification-api/internal/config"
	"github.com/AmirAbaris/notification-api/internal/db"
	"github.com/AmirAbaris/notification-api/internal/health"
	"github.com/AmirAbaris/notification-api/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found")
	}

	cfg := config.NewConfig()

	pool, err := db.NewPool(cfg.DBUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	healthHandler := health.NewHandler()
	userRepository := user.NewUserRepository(pool)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	r.GET("/health", healthHandler.GetHealth)
	r.POST("/users", userHandler.Create)

	r.Run(":" + cfg.Port)
}
