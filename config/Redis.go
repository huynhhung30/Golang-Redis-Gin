package config

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
    ctx = context.Background()
    RDB *redis.Client
)

func InitRedis() {
    RDB = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // điền nếu Redis có mật khẩu
        DB:       0,
    })
    pong, err := RDB.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Unable to connect to Redis: %v", err)
    }

    fmt.Println("Redis Response:", pong)
}
