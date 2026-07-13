package template

import (
	"context"

	"github.com/google/uuid"
)

type TemplateService struct {
	repo *TemplateRepository
}

func NewTemplateService(r *TemplateRepository) *TemplateService {
	return &TemplateService{repo: r}
}

func (s *TemplateService) Create(ctx context.Context, name string, subject, body string) (Template, error) {
	newTemplate, err := s.repo.Create(ctx, name, subject, body)
	if err != nil {
		return Template{}, err
	}

	return newTemplate, nil
}

func (s *TemplateService) Get(ctx context.Context, TemplateId uuid.UUID) (Template, error) {
	temp, err := s.repo.Get(ctx, TemplateId)
	if err != nil {
		return Template{}, err
	}

	return temp, nil
}
