package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	_ "trinity-app/cmd/docs"
	"trinity-app/config"
	errorsController "trinity-app/controllers/errors"
	"trinity-app/middlewares"
	"trinity-app/models"
	"trinity-app/routes"
	"trinity-app/utils/functions"

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
	}
}

var (
	redis_host = "127.0.0.1"
	redis_port = "6379"
	redis_uri  = fmt.Sprintf("redis://%s:%s/0", redis_host, redis_port)
)

func InitRedis() {
	opt, err := redis.ParseURL(redis_uri)
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opt)
	pong, err := rdb.Ping().Result()
	if err != nil {
		functions.ShowLog("Connect redis error", err.Error())
	}
	functions.ShowLog("Connect redis Success", pong, err)
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
