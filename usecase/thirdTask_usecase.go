package usecase

import (
	"context"

	"github.com/erdauletbatalov/tsarka/domain"
)

type counterUsecase struct {
	counterRepo domain.CounterRepository
}

func NewCounterUsecase(counter domain.CounterRepository) domain.CounterUsecase {
	return &counterUsecase{
		counterRepo: counter,
	}
}

func (count *counterUsecase) Add(ctx context.Context, num int) error {
	key, err := count.counterRepo.Get(ctx)
	if err != nil {
		return err
	}
	err = count.counterRepo.Set(ctx, key+num)
	if err != nil {
		return err
	}
	return nil
}

func (count *counterUsecase) Sub(ctx context.Context, num int) error {
	key, err := count.counterRepo.Get(ctx)
	if err != nil {
		return err
	}
	err = count.counterRepo.Set(ctx, key-num)
	if err != nil {
		return err
	}
	return nil
}

func (count *counterUsecase) Get(ctx context.Context) (int, error) {
	key, err := count.counterRepo.Get(ctx)
	if err != nil {
		return 0, err
	}
	return key, nil
}
