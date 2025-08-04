package main

import (
	_ "Golang-Redis-Gin/cmd/docs"
	"Golang-Redis-Gin/config"
	errorsController "Golang-Redis-Gin/internal/controllers/errors"

	"Golang-Redis-Gin/internal/middlewares"
	"Golang-Redis-Gin/internal/models"
	"Golang-Redis-Gin/internal/redis"
	"Golang-Redis-Gin/internal/routes"
	"Golang-Redis-Gin/internal/utils/functions"
	"net/http"
	"os"
	"time"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title 	Trinity App Tag Service API
// @version 1.0
// @description A tag service api
// @host 	localhost:5001
// @BasePath /api/v1/trinity
func main() {
	godotenv.Load("config/.env")
	// Khởi tạo kết nối Redis
	cfg := config.LoadConfig()
    rdb := redis.NewRedisClient(cfg.RedisAddress)
    defer rdb.Close()
	router := routes.ApplicationV1Router(rdb) 
	initialGinconfig(router)
	router.Use(middlewares.GinBodyLogMiddleware)
	router.Use(errorsController.Handler)
	go models.StartRpcServer()
	startServer(router)
}

func initialGinconfig(router *gin.Engine) {
	router.Use(limit.MaxAllowed(200))
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type, access-control-allow-origin, access-control-allow-headers, authorization  "},
	}))

	var err error
	config.DB, err = config.GormOpen()

	if err != nil {
		functions.ShowLog("Connect database error", err.Error())
	}else{
		functions.ShowLog("Connect database Success")
	}
}



func startServer(router http.Handler) {
	serverPort := os.Getenv("PORT")
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
