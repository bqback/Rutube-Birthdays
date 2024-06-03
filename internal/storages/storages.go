package storages

import (
	postgresql "birthdays/internal/storages/postgres"

	"github.com/jmoiron/sqlx"
)

type Storages struct {
	Auth IAuthStorage
	User IUserStorage
	Job  IJobStorage
}

func NewPostgresStorages(db *sqlx.DB) *Storages {
	return &Storages{
		Auth: postgresql.NewAuthStorage(db),
		User: postgresql.NewUserStorage(db),
		Job:  postgresql.NewJobStorage(db),
	}
}
