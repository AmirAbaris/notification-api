package queue

import (
	"context"

	"github.com/redis/go-redis/v9"
)

const RedisNotificatinKey = "notifications"

type Queue struct {
	client *redis.Client
}

func NewQueue(client *redis.Client) *Queue {
	return &Queue{client: client}
}

func (q *Queue) EnqueueNotification(
	ctx context.Context,
	notifictionID string,
) error {
	return q.client.LPush(
		ctx,
		RedisNotificatinKey,
		notifictionID,
	).Err()
}

func (q *Queue) Consume(
	ctx context.Context,
) (string, error) {
	result, err := q.client.BRPop(
		ctx,
		0,
		RedisNotificatinKey,
	).Result()

	if err != nil {
		return "", err
	}
	return result[1], nil
}
