package repository

import (
	"Golang-Redis-Gin/internal/cache"
	"Golang-Redis-Gin/internal/models"
	"Golang-Redis-Gin/internal/redis"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	redisLib "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)


type UserRepository interface {
    Save(ctx context.Context, user *models.UserModel) error
    FindByID(ctx context.Context, id string) (*models.UserModel, error)
}

type userRepository struct {
    gorm  *gorm.DB
    cache cache.Cache
}

func NewUserRepository(db *gorm.DB,r *redis.RedisCache) UserRepository {
    return &userRepository{
        gorm: db,
        cache: r,
    }
}

func (ur *userRepository) Save(ctx context.Context, user *models.UserModel) error {
    data, err := json.Marshal(user)
    if err != nil {
        return err
    }
    return ur.cache.Set(ctx, "user:"+strconv.Itoa(user.Id), data, 0)
}

func (ur *userRepository) FindByID(ctx context.Context, id string) (*models.UserModel, error) {
    cacheKey := fmt.Sprintf("user:%s", id) // id là string nên dùng %s

    // 1. Kiểm tra trong cache
    cached, err := ur.cache.Get(ctx, cacheKey)
    if err == nil {
        var user models.UserModel
        if jsonErr := json.Unmarshal([]byte(cached), &user); jsonErr == nil {
            fmt.Println("✅ Lấy từ cache Redis")
            return &user, nil
        }
    } else if err != redisLib.Nil {
        fmt.Println("⚠️ Redis error:", err)
    }
    
    // 2. Query from DB
    var user models.UserModel
	if err := ur.gorm.First(&user, id).Error; err != nil {
		return nil, err
	}
    // Lưu vào cache
	data, _ := json.Marshal(user)
	_ = ur.cache.Set(ctx, cacheKey, data, 5*time.Minute)
    return &user, nil
}
