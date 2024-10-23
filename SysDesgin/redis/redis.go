package redis

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	Client *redis.Client
}

func NewRedisClient() *Cache {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Could not connect to Redis:", err)
	}

	// Default set token
	_, error := rdb.Set(ctx, "tokens", 10, 0).Result()
	if error != nil {
		log.Fatal("Set token err", error)
	}
	_, error = rdb.Set(ctx, "lastFilled", time.Now(), 0).Result()
	if err != nil {
		log.Fatal("Set lastFilled err:", error)
	}

	return &Cache{
		Client: rdb,
	}
}
