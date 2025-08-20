package di

import (
	"Golang-Redis-Gin/internal/controllers"
	handler "Golang-Redis-Gin/internal/handlers"
	"Golang-Redis-Gin/internal/redis"
	"Golang-Redis-Gin/internal/repository"
	"Golang-Redis-Gin/internal/service"

	"gorm.io/gorm"
)

func InitUser(db *gorm.DB,r *redis.RedisCache) *handler.UserHandler{
    repo := repository.NewUserRepository(db,r)
    svc  := service.NewUserService(repo)
    ctrl := controllers.NewUserController(svc)
	return handler.NewUserHandler(ctrl)
}
