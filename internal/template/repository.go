package template

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TemplateRepository struct {
	db *pgxpool.Pool
}

type Template struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func NewTemplateRepository(db *pgxpool.Pool) *TemplateRepository {
	return &TemplateRepository{db: db}
}

func (r *TemplateRepository) Create(ctx context.Context, name, subject, body string) (Template, error) {
	var template Template
	err := r.db.QueryRow(ctx, `
	INSERT INTO templates(name, subject, body)
	VALUES($1, $2, $3)
	RETURNING id, name, subject, body, created_at;
	`,
		name,
		subject,
		body,
	).Scan(
		&template.ID,
		&template.Name,
		&template.Subject,
		&template.Body,
		&template.CreatedAt,
	)

	if err != nil {
		return Template{}, err
	}

	return template, nil
}

func (r *TemplateRepository) Get(ctx context.Context, templateId uuid.UUID) (Template, error) {
	var template Template
	err := r.db.QueryRow(ctx, `
	SELECT *
	FROM templates
	WHERE id = $1
	`,
		templateId,
	).Scan(
		&template.ID,
		&template.Name,
		&template.Subject,
		&template.Body,
		&template.CreatedAt,
	)

	if err != nil {
		return Template{}, err
	}

	return template, nil
}
