package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/AmirAbaris/notification-api/internal/config"
	"github.com/AmirAbaris/notification-api/internal/db"
	"github.com/AmirAbaris/notification-api/internal/notification"
	"github.com/AmirAbaris/notification-api/internal/queue"
	"github.com/AmirAbaris/notification-api/internal/redis"
	"github.com/AmirAbaris/notification-api/internal/template"
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

	redisClient := redis.NewClient(cfg.RedisUrl)
	queue := queue.NewQueue(redisClient)

	notificationRepository := notification.NewNotificationRepository(pool)

	templateRepository := template.NewTemplateRepository(pool)
	templateService := template.NewTemplateService(templateRepository)
	worker := NewWorker(queue, notificationRepository, templateService)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("all workers done")
			return

		default:
			err := worker.Process(ctx)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
