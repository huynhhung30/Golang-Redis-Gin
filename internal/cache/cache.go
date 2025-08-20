package cache

import (
	"context"
	"time"
)

// internal/cache/cache.go
type Cache interface {
    Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
    Get(ctx context.Context, key string) (string, error)
    Delete(ctx context.Context, key string) error
}
