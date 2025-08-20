package bootstrap

import (
	"Golang-Redis-Gin/config"
	"Golang-Redis-Gin/internal/database"
	"Golang-Redis-Gin/internal/redis"
	"Golang-Redis-Gin/internal/routes"
	"Golang-Redis-Gin/internal/utils/functions"
	"net/http"
	"os"
	"time"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Load config từ ENV
	cfg := config.LoadConfig()
	r := gin.Default()
	r.Use(limit.MaxAllowed(200))
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type, access-control-allow-origin, access-control-allow-headers, authorization  "},
	}))
	// Setup DB (gorm)
	functions.ShowLog("cfg", cfg.DB)
	var err error
	db, err := database.NewGormDB(cfg.DB)
	if err != nil {
		functions.ShowLog("Connect database error", err.Error())
	}else{
		functions.ShowLog("Connect database Success")
	}
	// init redis wrapper
	rdb := redis.NewRedisClient(cfg.Redis.Addr) 	

	// routes
	routes.ApplicationV1Router(db,r,rdb)

	return r
}

func StartServer(router http.Handler) {
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = "8001"
	}
	addr := ":" + serverPort

	s := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    18000 * time.Second,
		WriteTimeout:   18000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		functions.ShowLog("Start server error", err.Error())
	} else {
		functions.ShowLog("Start server success", s)
	}
}
