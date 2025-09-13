package service

import (
	"Golang-Redis-Gin/internal/cache"
	"Golang-Redis-Gin/internal/models"
	"Golang-Redis-Gin/internal/repository"
	"Golang-Redis-Gin/internal/utils"
	"Golang-Redis-Gin/internal/utils/functions"
	"context"
	"time"
)



type UserService interface {
    CreateUser(ctx context.Context, user *models.UserModel) (*models.UserModel, error)
    GetUser(ctx context.Context, id string) (*models.UserModel, error)
    LogInUser(ctx context.Context, user *models.UserModel) (*models.UserModel, error)
    GetUserByEmail(ctx context.Context, email string) (*models.UserModel, error)
    
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

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*models.UserModel, error) {
    user, err := s.repo.FindByEmail(ctx, email) // repo function to query DB
    if err != nil {
        return nil, utils.NewNotFound("user not found")
    }
    functions.ShowLog("emailemailemailemailemailemailemail", email)
    return user, nil
}
func (s *userService) LogInUser(ctx context.Context, user *models.UserModel) (*models.UserModel, error) {
    // Validate input
    if user.Email == "" {
        return nil, utils.NewBadRequest("email is required")
    }
    if user.Password == "" {
        return nil, utils.NewBadRequest("password is required")
    }
    if len(user.Password) < 6 {
        return nil, utils.NewBadRequest("password must be at least 6 characters")
    }
    
    // Check email conflict
    existing, _ := s.repo.FindByEmail(ctx, user.Email)
    if existing != nil {
        return nil, utils.NewConflict("email already exists")
    }

    // Hash password
    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        return nil, utils.NewInternal("cannot hash password")
    }
    user.Password = hashedPassword

    // Save user
    if err := s.repo.Save(ctx, user); err != nil {
        return nil, utils.NewInternal("failed to save user")
    }

    return user, nil
}
func (s *userService) CreateUser(ctx context.Context, user *models.UserModel) (*models.UserModel, error) {
    // Validate input
    if user.Email == "" {
        return nil, utils.NewBadRequest("email is required")
    }
    if user.Password == "" {
        return nil, utils.NewBadRequest("password is required")
    }
    if len(user.Password) < 6 {
        return nil, utils.NewBadRequest("password must be at least 6 characters")
    }
    
    // Check email conflict
    existing, _ := s.repo.FindByEmail(ctx, user.Email)
    if existing != nil {
        return nil, utils.NewConflict("email already exists")
    }

    // Hash password
    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        return nil, utils.NewInternal("cannot hash password")
    }
    user.Password = hashedPassword

    // Save user
    if err := s.repo.Save(ctx, user); err != nil {
        return nil, utils.NewInternal("failed to save user")
    }

    return user, nil
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