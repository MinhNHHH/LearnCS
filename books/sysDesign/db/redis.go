package db

import "github.com/redis/go-redis/v9"

func NewRedis() *redis.Client {
	options, err := redis.ParseURL("redis://localhost:6379")
	if err != nil {
		panic(err)
	}
	return redis.NewClient(options)
}
