package repository

import (
	"context"

	"github.com/erdauletbatalov/tsarka/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(ctx context.Context, user *domain.User) error {
	stmt := `
	INSERT INTO "user"(
		"first_name",
		"last_name"
	) VALUES ($1, $2)`

	_, err := u.db.Exec(ctx, stmt, user.FirstName, user.LastName)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Update(ctx context.Context, user *domain.User) error {
	stmt := `
					UPDATE "user" 
					SET "first_name" = $1, "last_name" = $2 
					WHERE "user_id" = $3`

	_, err := u.db.Exec(ctx, stmt, user.FirstName, user.LastName, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) Get(ctx context.Context, id int64) (*domain.User, error) {
	stmt := `SELECT * FROM "user" WHERE "user_id"=$1`
	user := domain.User{}
	err := u.db.QueryRow(ctx, stmt, id).Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) Delete(ctx context.Context, id int64) error {
	stmt := `DELETE FROM "user" WHERE "user_id"=$1`

	_, err := u.db.Exec(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}
