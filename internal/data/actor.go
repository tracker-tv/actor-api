package data

import "github.com/jackc/pgx/v5/pgtype"

type Actor struct {
	ID        int              `json:"id"`
	Name      string           `json:"name"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}
