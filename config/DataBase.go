package config

import (
	"Golang-Redis-Gin/utils/functions"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func GormOpen() (gormDB *gorm.DB, err error) {
	infodatabase := getDiverConn()
	gormDB, err = gorm.Open(mysql.Open(infodatabase.DriverConn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return 	
	}
functions.ShowLog("infodatabase.DriverConn",infodatabase.DriverConn)
	err = gormDB.Use(dbresolver.Register(dbresolver.Config{
		Replicas: []gorm.Dialector{mysql.Open(infodatabase.DriverConn)},
	}))
	return
}
