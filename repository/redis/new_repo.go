package redis

import (
	"context"
	"errors"

	"github.com/erdauletbatalov/tsarka/configs"
	"github.com/go-redis/redis/v8"
)

type Database struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

func NewRedisRepository(config *configs.Config) (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     configs.NewConfig().CacheAddr,
		Password: configs.NewConfig().CachePassword,
		DB:       0,
	})
	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	}
	return &Database{
		Client: client,
	}, nil
}
