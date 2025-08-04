package repository

import (
	"Golang-Redis-Gin/internal/cache"
	"Golang-Redis-Gin/internal/models"
	"context"
	"fmt"
	"time"
)

type UserRepository interface {
    FindByID(ctx context.Context, id string) (*models.UserModel, error)
}

type userRepo struct{}

func NewUserRepository() UserRepository {
    return &userRepo{}
}

func (r *userRepo) FindByID(ctx context.Context, id string) (*models.UserModel, error) {
    key := "user:" + id
    var user models.UserModel

    // 1. Redis cache
    if err := cache.GetCache(key, &user); err == nil {
        fmt.Println("✅ Lấy từ Redis")
        return &user, nil
    }

    // 2. Fake DB call
    user = models.UserModel{
        Id:user.Id,
        LastName:  "User " + id,
        Email: "user" + id + "@example.com",
    }

    // 3. Cache lại
    cache.SetCache(key, user, 5*time.Minute)
    fmt.Println("❌ Không có cache, tạo mới")

    return &user, nil
}
