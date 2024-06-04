package storages

import (
	"birthdays/internal/pkg/entities"
	"context"
)

type IJobStorage interface {
	GetAll(context.Context) ([]*entities.Job, error)
	Add(context.Context, entities.Job) error
	GetRecipientEmails(context.Context, uint64) ([]string, error)
}
