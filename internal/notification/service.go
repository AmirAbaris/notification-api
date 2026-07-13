package notification

import (
	"context"

	"github.com/google/uuid"
)

type NotificationService struct {
	repo *NotificationRepository
}

func NewNotificationService(r *NotificationRepository) *NotificationService {
	return &NotificationService{repo: r}
}

func (s *NotificationService) Create(ctx context.Context, userID, templateID uuid.UUID, status string) (Notification, error) {
	notification, err := s.repo.Create(ctx, userID, templateID, status)
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
