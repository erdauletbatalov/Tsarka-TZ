package domain

import "context"

type CounterUsecase interface {
	Add(ctx context.Context, num int) error
	Sub(ctx context.Context, num int) error
	Get(ctx context.Context) (int, error)
}

type CounterRepository interface {
	Set(ctx context.Context, num int) error
	Get(ctx context.Context) (int, error)
}
