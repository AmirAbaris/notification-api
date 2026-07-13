package user

import (
	"context"

	"github.com/google/uuid"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(r *UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Create(ctx context.Context, email, name string) (User, error) {
	user, err := s.repo.Create(ctx, name, email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *UserService) Get(ctx context.Context, userId uuid.UUID) (User, error) {
	user, err := s.repo.Get(ctx, userId)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
