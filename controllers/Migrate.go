package controllers

import (
	"Golang-Redis-Gin/config"
	"Golang-Redis-Gin/models"
	"Golang-Redis-Gin/utils/functions"
)

// Auto migrate
func Migrate() {
	config.DB.AutoMigrate(
	models.UserModel{},
	models.CouponModel{},
	)
	functions.ShowLog("MigrateModel", "Success")
}
