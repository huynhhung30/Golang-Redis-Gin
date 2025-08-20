package handler

import (
	"Golang-Redis-Gin/internal/redis"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetHandler(r *redis.RedisCache) gin.HandlerFunc {
    return func(c *gin.Context) {
        key := c.Query("key")
        value := c.Query("value")

        if key == "" || value == "" {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Lost key or value in query string",
            })
            return
        }

        // TTL: bạn có thể hardcode, hoặc lấy từ query nếu muốn
        ttl := 10 * time.Second

        err := r.Set(c.Request.Context(), key, value, ttl)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "error save redis: " + err.Error(),
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "OK saved Redis succes",
            "key":     key,
            "value":   value,
            "ttl":     ttl.Seconds(),
        })
    }
}
