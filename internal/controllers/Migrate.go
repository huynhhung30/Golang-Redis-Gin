package controllers

import (
	"Golang-Redis-Gin/internal/models"
	"Golang-Redis-Gin/internal/utils/functions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Auto migrate
func MigrateTable(db *gorm.DB,c *gin.Context) {
	db.Debug().AutoMigrate(
		&models.UserModel{},
	// models.CouponModel{},
	)
	functions.ShowLog("MigrateModel", "Success")
}
