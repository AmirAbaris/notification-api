package user

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRespository() *Repository {
	return &Repository{}
}

func (r *Repository) Create(ctx context.Context, email, name string) error {
	_, err := r.db.Exec(ctx, `
	INSERT INTO users (email, name)
	VALUES($1, $2)`,
		email, name,
	)

	return err
}
