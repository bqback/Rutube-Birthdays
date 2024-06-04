package postgresql

import (
	"birthdays/internal/apperrors"
	"birthdays/internal/pkg/dto"
	"birthdays/internal/pkg/entities"
	"context"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/go-chi/httplog/v2"
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

func (s *PgUserStorage) GetAll(ctx context.Context) ([]*entities.User, error) {
	httplog.LogEntrySetField(ctx, dto.StepKey, slog.StringValue(step))
	httplog.LogEntrySetField(ctx, dto.FuncKey, slog.StringValue("GetAll"))
	oplog := httplog.LogEntry(ctx)

	query, args, err := squirrel.
		Select(userFullSelectFields...).
		From(authTable).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		oplog.Debug("Failed to build query", "err", err.Error())
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	oplog.Debug("Query built")

	var users []*entities.User
	if err := s.db.Select(&users, query, args...); err != nil {
		oplog.Debug("User select failed", "err", err.Error())
		return nil, apperrors.ErrUserNotSelected
	}
	oplog.Debug("Query executed")

	if len(users) == 0 {
		return nil, apperrors.ErrEmptyResult
	}

	return users, nil
}
