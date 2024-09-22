package data

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Actor struct {
	ID        int              `json:"id"`
	Name      string           `json:"name"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type ActorModel interface {
	GetActors() ([]*Actor, error)
}

type PgActorModel struct {
	DB *pgx.Conn
}

func (m PgActorModel) GetActors() ([]*Actor, error) {
	query := `SELECT id, name, created_at, updated_at FROM actors`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := m.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	actors := make([]*Actor, 0)
	for rows.Next() {
		var actor Actor
		err := rows.Scan(&actor.ID, &actor.Name, &actor.CreatedAt, &actor.UpdatedAt)
		if err != nil {
			return nil, err
		}
		actors = append(actors, &actor)
	}

	return actors, nil
}
