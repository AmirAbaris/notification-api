package main

import (
	"context"
	"fmt"

	"github.com/AmirAbaris/notification-api/internal/notification"
	"github.com/AmirAbaris/notification-api/internal/queue"
	"github.com/AmirAbaris/notification-api/internal/renderer"
	"github.com/AmirAbaris/notification-api/internal/template"
	"github.com/google/uuid"
)

type Worker struct {
	queue            *queue.Queue
	notificationRepo *notification.NotificationRepository
	templateService  *template.TemplateService
}

func NewWorker(
	queue *queue.Queue,
	notificationRepo *notification.NotificationRepository,
	templateService *template.TemplateService,

) *Worker {
	return &Worker{
		queue:            queue,
		notificationRepo: notificationRepo,
		templateService:  templateService,
	}
}

func (w *Worker) Process(ctx context.Context) error {
	// notifs, err := w.notificationRepo.ClaimPending(ctx, 10)
	notificationID, err := w.queue.Consume(ctx)
	if err != nil {
		return err
	}

	notif, err := w.notificationRepo.Get(ctx, uuid.MustParse(notificationID))
	if err != nil {
		return err
	}

	tmpl, err := w.templateService.Get(ctx, notif.TemplateID)
	if err != nil {
		return err
	}
	message := renderer.Render(tmpl.Body, notif.Data)
	// fake sender
	fmt.Println(message)
	_, err = w.notificationRepo.UpdateToSent(ctx, notif.ID)
	if err != nil {
		return err
	}

	return nil
}
