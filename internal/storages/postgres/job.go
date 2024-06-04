package postgresql

import (
	"birthdays/internal/pkg/entities"
	"context"

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

func (s *PgJobStorage) GetAll(ctx context.Context) ([]*entities.Job, error) {
	return nil, nil
}

func (s *PgJobStorage) Add(ctx context.Context, job entities.Job) error {
	return nil
}

func (s *PgJobStorage) GetRecipientEmails(ctx context.Context, sourceID uint64) ([]string, error) {
	return []string{""}, nil
}
