package user

import (
	"context"
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
