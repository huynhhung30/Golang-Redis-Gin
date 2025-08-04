package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
    client *redis.Client
    ctx    context.Context
}

var (
    RDB *redis.Client
)
func NewRedisClient(addr string) *RedisClient {
    ctx := context.Background()
    rdb := redis.NewClient(&redis.Options{
        Addr: addr,
        DB:   0,
    })

    return &RedisClient{
        client: rdb,
        ctx:    ctx,
    }
}

func (r *RedisClient) GetCtx() context.Context {
    return r.ctx
}

func (r *RedisClient) Client() *redis.Client {
    return r.client
}

func (r *RedisClient) Close() error {
    return r.client.Close()
}
