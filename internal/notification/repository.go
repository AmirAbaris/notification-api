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
	ID         uuid.UUID         `json:"id"`
	UserID     uuid.UUID         `json:"user_id"`
	TemplateID uuid.UUID         `json:"template_id"`
	Data       map[string]string `json:"data"`
	Status     string            `json:"status"`
	CreatedAt  time.Time         `json:"created_at"`
}

func NewNotificationRepository(db *pgxpool.Pool) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (r *NotificationRepository) Create(ctx context.Context, userID, templateID uuid.UUID, data map[string]string) (Notification, error) {
	var notification Notification

	err := r.db.QueryRow(ctx, `
	INSERT INTO notifications (user_id, template_id, data)
	VALUES($1, $2, $3)
	RETURNING id, user_id, template_id, status, data, created_at;
	`,
		userID,
		templateID,
		data,
	).Scan(
		&notification.ID,
		&notification.UserID,
		&notification.TemplateID,
		&notification.Status,
		&notification.Data,
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
	SELECT id, user_id, template_id, data, status, created_at
	FROM notifications
	WHERE id = $1
	`,
		notificationID,
	).Scan(
		&notification.ID,
		&notification.UserID,
		&notification.TemplateID,
		&notification.Data,
		&notification.Status,
		&notification.CreatedAt,
	)

	if err != nil {
		return Notification{}, err
	}

	return notification, nil
}

func (r *NotificationRepository) GetPendings(ctx context.Context, limit int) ([]Notification, error) {
	rows, err := r.db.Query(ctx, `
	SELECT id, user_id, template_id, data, status, created_at
	FROM notifications
	WHERE status = 'pending'
	LIMIT $1
	`,
		limit,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var notifications []Notification

	for rows.Next() {
		var notification Notification

		err := rows.Scan(
			&notification.ID,
			&notification.UserID,
			&notification.TemplateID,
			&notification.Data,
			&notification.Status,
			&notification.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		notifications = append(notifications, notification)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}

func (r *NotificationRepository) UpdateToSent(ctx context.Context, notificationID uuid.UUID) (Notification, error) {
	var notification Notification

	err := r.db.QueryRow(ctx, `
	UPDATE notifications
	SET status = 'sent'
	WHERE id = $1
	RETURNING id, user_id, template_id, data, status, created_at;
	`,
		notificationID,
	).Scan(
		&notification.ID,
		&notification.UserID,
		&notification.TemplateID,
		&notification.Data,
		&notification.Status,
		&notification.CreatedAt,
	)

	if err != nil {
		return Notification{}, err
	}

	return notification, nil
}

func (r *NotificationRepository) ClaimPending(ctx context.Context, limit int) ([]Notification, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	rows, err := r.db.Query(ctx, `
	SELECT id, user_id, template_id, data, status, created_at
	FROM notifications
	WHERE status = 'pending'
	FOR UPDATE SKIP LOCKED
	LIMIT $1
	`,
		limit,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var notifications []Notification

	for rows.Next() {
		var notification Notification

		err := rows.Scan(
			&notification.ID,
			&notification.UserID,
			&notification.TemplateID,
			&notification.Data,
			&notification.Status,
			&notification.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		notifications = append(notifications, notification)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Claim them
	for _, notif := range notifications {
		_, err = tx.Exec(ctx, `
				UPDATE notifications
				SET status = 'processing'
				WHERE id = $1
			`,
			notif.ID,
		)
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}
