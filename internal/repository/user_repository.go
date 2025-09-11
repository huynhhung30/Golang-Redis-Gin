package repository

import (
	"Golang-Redis-Gin/internal/models"
	"context"

	"gorm.io/gorm"
)


type UserRepository interface {
    Save(ctx context.Context, user *models.UserModel) error
    FindByID(ctx context.Context, id string) (*models.UserModel, error)
    FindByEmail(ctx context.Context, email string) (*models.UserModel, error)

}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

// Save user vào DB
func (r *userRepository) Save(ctx context.Context, user *models.UserModel) error {
    return r.db.WithContext(ctx).Debug().Create(user).Error
}

// Query theo ID từ DB
func (r *userRepository) FindByID(ctx context.Context, id string) (*models.UserModel, error) {
    var user models.UserModel
    if err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*models.UserModel, error) {
    var user models.UserModel
    if err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

