package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/tracker-tv/actor-api/internal/config"
)

func OpenDB(cfg config.Config) (*pgx.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := pgx.Connect(context.Background(), cfg.DB)
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
