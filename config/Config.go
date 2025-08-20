package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type RedisConfig struct {
	Addr string
}

type AppConfig struct {
	DB    DatabaseConfig
	Redis RedisConfig
	Port  string
}

func LoadConfig() *AppConfig {
		// Load env file
		err := godotenv.Load("config/.env")
		if err != nil {
			log.Fatal("❌ Error loading .env file")
		}
	return &AppConfig{
		DB: DatabaseConfig{
			Driver:   "mysql",
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		Redis: RedisConfig{
			Addr: os.Getenv("REDIS_ADDR"),
		},
		Port: os.Getenv("PORT"),
	}
}




// package config

// import "os"



// type Config struct {
//     RedisAddress string
// }

// func LoadConfig() *Config {
//     return &Config{
//         RedisAddress: getEnv("REDIS_ADDR", "localhost:6379"),
//     }
// }

// func getEnv(key, fallback string) string {
//     if value, ok := os.LookupEnv(key); ok {
//         return value
//     }
//     return fallback
// }