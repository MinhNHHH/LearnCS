package database

import (
	"fmt"

	"github.com/go-redis/redis"
)

func RedisClient() *redis.Client {
	// db, _err := strconv.ParseInt(os.Getenv("REDIS_DB_NAME"), 10, 64)
	// if _err != nil {
	// 	db = 0
	// }
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}
