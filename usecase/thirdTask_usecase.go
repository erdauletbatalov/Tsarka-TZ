package usecase

import (
	"context"
	"strconv"

	"github.com/erdauletbatalov/tsarka/domain"
)

const counter = "counter"

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
	numFromStr, err := strconv.Atoi(key)
	if err != nil {
		return err
	}
	err = count.counterRepo.Set(ctx, numFromStr+num)
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
	numFromStr, err := strconv.Atoi(key)
	if err != nil {
		return err
	}
	err = count.counterRepo.Set(ctx, numFromStr-num)
	if err != nil {
		return err
	}
	return nil
}

func (count *counterUsecase) Get(ctx context.Context) (string, error) {
	key, err := count.counterRepo.Get(ctx)
	if err != nil {
		return "", err
	}
	return key, nil
}
