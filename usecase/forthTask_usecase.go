package usecase

import (
	"context"
	"strings"

	"github.com/erdauletbatalov/tsarka/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (u *userUsecase) Create(ctx context.Context, user *domain.User) error {
	if strings.TrimSpace(user.FirstName) == "" || strings.TrimSpace(user.LastName) == "" {
		return domain.ErrNotValid
	}

	return u.userRepo.Create(ctx, user)
}

func (u *userUsecase) Update(ctx context.Context, user *domain.User) error {

	err := u.userRepo.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) Get(ctx context.Context, id int64) (*domain.User, error) {
	user, err := u.userRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) Delete(ctx context.Context, id int64) error {
	err := u.userRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
