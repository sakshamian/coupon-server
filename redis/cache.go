package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

type SetRedis struct {
	Key  string
	Data interface{}
	Exp  time.Duration
}

// Global variables for Redis clients
var (
	RedisClient *redis.Client
	ctx         = context.Background()
)

// Initialize redis client
func Connect() (*redis.Client, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found or there was an error loading it: %v", err)
	}

	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "redis:6379"
	}
	redisPassword := os.Getenv("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
	})
	pong, err := client.Ping(&gin.Context{}).Result()
	fmt.Println("[redis]", pong, " ", err)

	if err != nil {
		log.Println("[redis] not initialize with error: ", err)
		return nil, err
	}
	RedisClient = client
	return RedisClient, nil
}

func SetInRedis(set_data *SetRedis) error {
	json, err := json.Marshal(set_data.Data)
	if err != nil {
		return err
	}

	RedisClient.Set(ctx, set_data.Key, json, set_data.Exp)

	if os.Getenv("DEBUG") == "true" {
		fmt.Printf("Key: %s set into Redis\n", set_data.Key)
	}

	return nil
}

func GetFromRedis(key string) (*string, error) {
	val, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Println("Message successfully got from Redis")
	}
	return &val, nil
}

func DeleteFromRedis(key string) (*int64, error) {
	val, err := RedisClient.Del(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Println("Message successfully deleted in Redis")
	}
	return &val, nil
}

func DeletePatternKeys(pattern string) error {
	iter := RedisClient.Scan(ctx, 0, pattern, 0).Iterator()
	for iter.Next(ctx) {
		if _, err := RedisClient.Del(ctx, iter.Val()).Result(); err != nil {
			return err
		}
	}
	if err := iter.Err(); err != nil {
		return err
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Printf("Redis: Successfully deleted keys matching pattern: %s\n", pattern)
	}

	return nil
}
