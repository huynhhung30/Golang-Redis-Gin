package main

import (
	_ "Golang-Redis-Gin/cmd/docs"
	"Golang-Redis-Gin/config"
	errorsController "Golang-Redis-Gin/controllers/errors"
	"Golang-Redis-Gin/middlewares"
	"Golang-Redis-Gin/models"
	"Golang-Redis-Gin/routes"
	"Golang-Redis-Gin/utils/functions"
	"net/http"
	"os"
	"time"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

// @title 	Trinity App Tag Service API
// @version 1.0
// @description A tag service api
// @host 	localhost:5001
// @BasePath /api/v1/trinity
func main() {
	godotenv.Load(".env")
	router := gin.Default()
	initialGinConfig(router)
	InitRedis()
	router.Use(middlewares.GinBodyLogMiddleware)
	router.Use(errorsController.Handler)
	routes.ApplicationV1Router(router)
	// controllers.Migrate()
	go models.StartRpcServer()
	startServer(router)
}

func initialGinConfig(router *gin.Engine) {
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
		functions.ShowLog("Connect Database Success")
	}
}

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := rdb.Ping().Result()
	if err != nil {
		functions.ShowLog("Connect redis error", err.Error())
	}else{
		functions.ShowLog("Connect redis Success", pong, err)
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
