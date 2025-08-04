package config

import "os"

const SECRET = "LIFE_ON_BE"
const SECRET_ADMIN = "SECRET_ADMIN"
const EXP_HOURS = 1000
const EXP_HOURS_ADMIN = 1000

type Config struct {
    RedisAddress string
}

func LoadConfig() *Config {
    return &Config{
        RedisAddress: getEnv("REDIS_ADDR", "localhost:6379"),
    }
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}