package services

import (
	"context"

	"github.com/n340r/backend-notes/toptal/internal/app/domain"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (s UserService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	return s.repo.CreateUser(ctx, user)
}

func (s UserService) GetUser(ctx context.Context, username string) (domain.User, error) {
	return s.repo.GetUser(ctx, username)
}

func (s UserService) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	return s.repo.GetUserByID(ctx, id)
}
