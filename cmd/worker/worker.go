package worker

import (
	"context"
	"fmt"
	"log"

	"github.com/AmirAbaris/notification-api/internal/notification"
	"github.com/AmirAbaris/notification-api/internal/renderer"
	"github.com/AmirAbaris/notification-api/internal/template"
)

type Worker struct {
	notificationRepo *notification.NotificationRepository
	templateService  *template.TemplateService
}

func NewWorker(
	notificationRepo *notification.NotificationRepository,
	templateService *template.TemplateService,

) *Worker {
	return &Worker{
		notificationRepo: notificationRepo,
		templateService:  templateService,
	}
}

func (w *Worker) Process(ctx context.Context) error {
	notifs, err := w.notificationRepo.ClaimPending(ctx, 10)
	if err != nil {
		return err
	}

	for _, notif := range notifs {
		tmpl, err := w.templateService.Get(ctx, notif.TemplateID)
		if err != nil {
			log.Println(err)
			continue
		}
		message := renderer.Render(tmpl.Body, notif.Data)
		// fake sender
		fmt.Println(message)
		_, err = w.notificationRepo.UpdateToSent(ctx, notif.ID)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	return nil
}
