package service

import (
	"Golang-Redis-Gin/internal/cache"
	"Golang-Redis-Gin/internal/models"
	"Golang-Redis-Gin/internal/repository"
	"Golang-Redis-Gin/internal/utils/functions"
	"context"
	"strconv"
	"time"
)



type UserService interface {
    CreateUser(ctx context.Context, user *models.UserModel) error
    GetUser(ctx context.Context, id string) (*models.UserModel, error)
}

type userService struct {
    repo repository.UserRepository
    cache cache.Cache
}

func NewUserService(repo repository.UserRepository, cache cache.Cache) UserService {
    return &userService{
        repo:  repo,
        cache: cache,
    }
}


func (s *userService) CreateUser(ctx context.Context, user *models.UserModel) error {
    // Business logic (ví dụ: validate)
    if strconv.Itoa(user.Id) == "" || user.FirstName == ""|| user.LastName == "" {
        return nil
    }
    return s.repo.Save(ctx, user)
}

func (s *userService) GetUser(ctx context.Context, id string) (*models.UserModel, error) {
    key := "user:" + id
    var cachedUser models.UserModel

    // 1. Lấy từ Redis trước
    if err := s.cache.Get(ctx, key, &cachedUser); err == nil {
        functions.ShowLog("Cache hit:", key)

        return &cachedUser, nil
    }

    // 2. Nếu cache miss → DB
    functions.ShowLog("Cache miss -> DB:", key)
    user, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }

    // 3. Lưu lại vào Redis
    _ = s.cache.Set(ctx, key, user, time.Hour)

    return user, nil
}