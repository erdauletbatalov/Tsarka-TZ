package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/erdauletbatalov/tsarka/configs"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgresRepository(config *configs.Config) (*pgxpool.Pool, error) {
	DSN := fmt.Sprintf("postgres://%s:%s@%s%s/%s", config.UserDB, config.PasswordDB, config.HostDB, config.PortDB, config.NameDB)

	cfg, err := pgxpool.ParseConfig(DSN)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", err)
	}

	cfg.MaxConns = 5 // TODO 25
	cfg.MaxConnLifetime = 5 * time.Minute

	db, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", err)
	}

	return db, nil
}
