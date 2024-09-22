package data

import "github.com/jackc/pgx/v5"

type Models struct {
	ActorModel ActorModel
}

func NewModels(db *pgx.Conn) Models {
	return Models{
		ActorModel: PgActorModel{DB: db},
	}
}
