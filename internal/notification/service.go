package notification

import (
	"context"

	"github.com/AmirAbaris/notification-api/internal/queue"
	"github.com/google/uuid"
)

type NotificationService struct {
	repo *NotificationRepository
	q    *queue.Queue
}

func NewNotificationService(r *NotificationRepository, queue *queue.Queue) *NotificationService {
	return &NotificationService{repo: r, q: queue}
}

func (s *NotificationService) Create(ctx context.Context, userID, templateID uuid.UUID, data map[string]string) (Notification, error) {
	notification, err := s.repo.Create(ctx, userID, templateID, data)
	if err != nil {
		return Notification{}, err
	}

	err = s.q.EnqueueNotification(ctx, notification.ID.String())

	if err != nil {
		return Notification{}, err
	}

	return notification, nil
}

func (s *NotificationService) Get(ctx context.Context, notificationID uuid.UUID) (Notification, error) {
	notification, err := s.repo.Get(ctx, notificationID)
	if err != nil {
		return Notification{}, err
	}

	return notification, nil
}
