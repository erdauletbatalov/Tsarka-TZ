package redis

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

const counter = "key"

type Database struct {
	Client redis.Conn
}

func NewRedisRepository(password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "RedisTexzadanie2022",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Println(err)
	}
	err = client.Set("key", "0", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pong, err)
	return client
}
