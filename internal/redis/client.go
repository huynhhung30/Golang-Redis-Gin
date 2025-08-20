package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(addr string) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	
	// kiểm tra kết nối
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("❌ Cannot connect to Redis: %v", err)
	}

	log.Println("✅ Connected to Redis")

	return &RedisCache{client: rdb}
}