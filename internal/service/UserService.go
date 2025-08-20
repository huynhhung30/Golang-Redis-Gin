package service

import (
	"Golang-Redis-Gin/internal/models"
	"Golang-Redis-Gin/internal/repository"
	"context"
	"strconv"
)



type UserService interface {
    CreateUser(ctx context.Context, user *models.UserModel) error
    GetUser(ctx context.Context, id string) (*models.UserModel, error)
}

type userService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, user *models.UserModel) error {
    // Business logic (ví dụ: validate)
    if strconv.Itoa(user.Id) == "" || user.FirstName == ""|| user.LastName == "" {
        return nil
    }
    return s.repo.Save(ctx, user)
}

func (s *userService) GetUser(ctx context.Context, id string) (*models.UserModel, error) {
    return s.repo.FindByID(ctx, id)
}