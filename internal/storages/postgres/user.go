package postgresql

import (
	"github.com/jmoiron/sqlx"
)

type PgUserStorage struct {
	db *sqlx.DB
}

func NewUserStorage(db *sqlx.DB) *PgUserStorage {
	return &PgUserStorage{
		db: db,
	}
}
