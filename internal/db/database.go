package database

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var (
	Client *redis.Client
	DbCtx  = context.Background()
)

func Init(address string) {
	Client = redis.NewClient(&redis.Options{
		Addr:     getEnv("REDIS_ADDR", address),
		Password: "",
		DB:       0,
	})

	if err := ping(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Printf("Redis client running on %s\n", address)
}

func ping() error {
	_, err := Client.Ping(DbCtx).Result()
	return err
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
