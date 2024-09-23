package testutils

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func SetupDB(dsn string) (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	seeds := []string{"./../../testdata/actors.sql"}
	for _, seed := range seeds {
		query, err := os.ReadFile(seed)
		if err != nil {
			return nil, err
		}
		_, err = db.Exec(context.Background(), string(query))
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
