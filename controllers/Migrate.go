package controllers

import (
	"Golang-Redis-Gin/config"
	"Golang-Redis-Gin/models"
	"Golang-Redis-Gin/utils/functions"

	"github.com/gin-gonic/gin"
)

// Auto migrate
func MigrateTable(c *gin.Context) {
	config.DB.Debug().AutoMigrate(
		&models.UserModel{},
	// models.CouponModel{},
	)
	functions.ShowLog("MigrateModel", "Success")
}
