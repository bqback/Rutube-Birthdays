package postgresql

import (
	"github.com/jmoiron/sqlx"
)

type PgJobStorage struct {
	db *sqlx.DB
}

func NewJobStorage(db *sqlx.DB) *PgJobStorage {
	return &PgJobStorage{
		db: db,
	}
}
