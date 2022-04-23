package repository

import (
	"context"
	"time"

	"github.com/erdauletbatalov/tsarka/domain"
	redisDb "github.com/erdauletbatalov/tsarka/repository/redis"
)

const counter = "counter"

type counterRepository struct {
	db *redisDb.Database
}

func NewCounterRepository(db *redisDb.Database) domain.CounterRepository {
	return &counterRepository{
		db: db,
	}
}

func (coun *counterRepository) Set(ctx context.Context, num int) error {
	pipe := coun.db.Client.TxPipeline()
	new := pipe.Set(ctx, counter, num, time.Duration(time.Minute*30))
	if new.Err() != nil {
		return new.Err()
	}
	return nil
}

func (coun *counterRepository) Get(ctx context.Context) (string, error) {
	pipe := coun.db.Client.TxPipeline()
	result := pipe.Get(ctx, counter)
	if result.Err() != nil {
		return "", result.Err()
	}

	return result.String(), nil
}
