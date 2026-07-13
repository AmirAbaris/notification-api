package notification

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NotificationRepository struct {
	db *pgxpool.Pool
}

type Notification struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	TemplateID uuid.UUID `json:"template_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewNotificationRepository(db *pgxpool.Pool) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (r *NotificationRepository) Create(ctx context.Context, userID, templateID uuid.UUID, status string) (Notification, error) {
	var notification Notification

	err := r.db.QueryRow(ctx, `
	INSERT INTO notifications (user_id, template_id, status)
	VALUES($1, $2, $3)
	RETURNING id, user_id, template_id, status, created_at;
	`,
		userID,
		templateID,
		status,
	).Scan(
		&notification.ID,
		&notification.UserID,
		&notification.TemplateID,
		&notification.Status,
		&notification.CreatedAt,
	)

	if err != nil {
		return Notification{}, err
	}

	return notification, nil
}

func (r *NotificationRepository) Get(ctx context.Context, notificationID uuid.UUID) (Notification, error) {
	var notification Notification

	err := r.db.QueryRow(ctx, `
	SELECT id, user_id, template_id, status, created_at
	FROM notifications
	WHERE id = $1
	`,
		notificationID,
	).Scan(
		&notification.ID,
		&notification.UserID,
		&notification.TemplateID,
		&notification.Status,
		&notification.CreatedAt,
	)

	if err != nil {
		return Notification{}, err
	}

	return notification, nil
}
