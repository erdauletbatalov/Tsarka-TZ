package repository

import (
	"context"

	"github.com/erdauletbatalov/tsarka/domain"
	"github.com/go-redis/redis"
)

const counter = "key"

type counterRepository struct {
	db *redis.Client
}

func NewCounterRepository(db *redis.Client) domain.CounterRepository {
	return &counterRepository{
		db: db,
	}
}

func (coun *counterRepository) Set(ctx context.Context, num int) error {
	err := coun.db.Set(counter, num, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (coun *counterRepository) Get(ctx context.Context) (int, error) {
	new, err := coun.db.Get(counter).Int()
	if err != nil {
		return 0, err
	}
	return new, nil
}
