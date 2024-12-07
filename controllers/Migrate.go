package controllers

import (
	"trinity-app/config"
	"trinity-app/models"
	"trinity-app/utils/functions"
)

// Auto migrate
func Migrate() {
	config.DB.AutoMigrate(
	models.UserModel{},
	models.CouponModel{},
	)
	functions.ShowLog("MigrateModel", "Success")
}
