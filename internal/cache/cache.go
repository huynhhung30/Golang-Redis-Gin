package cache

import (
	"Golang-Redis-Gin/internal/redis"
	"context"
	"encoding/json"
	"time"
)

func SetCache(key string, data any, ttl time.Duration) error {
    b, _ := json.Marshal(data)
    return redis.RDB.Set(context.Background(), key, b, ttl).Err()
}

func GetCache(key string, result any) error {
    val, err := redis.RDB.Get(context.Background(), key).Result()
    if err != nil {
        return err
    }
    return json.Unmarshal([]byte(val), result)
}
