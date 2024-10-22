package redis

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
}

func NewRedisClient() *Redis {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Could not connect to Redis: %v", err))
	}
	return &Redis{
		client: rdb,
	}
}
