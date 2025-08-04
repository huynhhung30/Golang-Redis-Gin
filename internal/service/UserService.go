package service

import (
	"Golang-Redis-Gin/internal/models"
	"Golang-Redis-Gin/internal/repository"
	"context"
)

type UserService interface {
    GetProfile(ctx context.Context, id string) (*models.UserModel, error)
}

type userService struct {
    repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
    return &userService{repo: r}
}

func (s *userService) GetProfile(ctx context.Context, id string) (*models.UserModel, error) {
    return s.repo.FindByID(ctx, id)
}
