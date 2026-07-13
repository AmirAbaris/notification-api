package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, email, name string) (User, error) {
	var user User
	err := r.db.QueryRow(ctx, `
	INSERT INTO users (email, name)
	VALUES($1, $2)
	RETURNING id, email, name, created_at;
	`,
		email, name,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.CreatedAt,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *UserRepository) Get(ctx context.Context, userId uuid.UUID) (User, error) {
	var user User
	err := r.db.QueryRow(ctx, `
	SELECT *
	FROM users
	WHERE id = $1
	`,
		userId,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.CreatedAt,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}
