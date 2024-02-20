package redis

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

const (
	REDIS_HOST = "REDIS_HOST"
	REDIS_PORT = "REDIS_PORT"
)

func GetRedisClient() *redis.Client {
	redisHost := os.Getenv(REDIS_HOST)
	redisPort := os.Getenv(REDIS_PORT)

	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "",
		DB:       0,
	})
}
