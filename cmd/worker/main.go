package worker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AmirAbaris/notification-api/internal/config"
	"github.com/AmirAbaris/notification-api/internal/db"
	"github.com/AmirAbaris/notification-api/internal/notification"
	"github.com/AmirAbaris/notification-api/internal/renderer"
	"github.com/AmirAbaris/notification-api/internal/template"
	"github.com/joho/godotenv"
)

func main() {
	counter := 0
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

	notificationRepository := notification.NewNotificationRepository(pool)

	templateRepository := template.NewTemplateRepository(pool)
	templateService := template.NewTemplateService(templateRepository)

	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)

	for {
		select {
		case <-ticker.C:
			notifs, err := notificationRepository.GetPendings(context.Background(), counter)

			for _, notif := range notifs {
				template, _ := templateService.Get(context.Background(), notif.TemplateID)
				_ = renderer.Render(template.Body, notif.Data)

				// fake sender
				fmt.Println("sent")

				// update
				updatedNotif, err := notificationRepository.UpdateToSent(context.Background(), notif.ID)
				if err != nil {
					log.Fatal(err)
					return
				}
				fmt.Println(updatedNotif)
			}
			if err != nil {
				log.Fatal(err)
				return
			}
			counter = counter + 10

		case <-done:
			fmt.Println("all workers done")
		}
	}
}
