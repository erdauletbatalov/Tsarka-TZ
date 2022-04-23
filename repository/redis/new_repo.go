package redis

import (
	"context"
	"errors"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type Database struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

func NewRedisRepository(addr string) (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "RedisTexzadanie2022",
		DB:       0,
	})
	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	}

	pipe := client.TxPipeline()
	pipe.Set(Ctx, "counter", "0", time.Duration(time.Minute*30))

	return &Database{
		Client: client,
	}, nil
}
