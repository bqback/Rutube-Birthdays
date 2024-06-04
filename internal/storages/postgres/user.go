package postgresql

import (
	"birthdays/internal/apperrors"
	"birthdays/internal/pkg/dto"
	"birthdays/internal/pkg/entities"
	"context"
	"errors"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/go-chi/httplog/v2"
	"github.com/jackc/pgx/v5/pgconn"
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
		From(userTable).
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

func (s *PgUserStorage) Subscribe(ctx context.Context, source uint64, id uint64) error {
	httplog.LogEntrySetField(ctx, dto.StepKey, slog.StringValue(step))
	httplog.LogEntrySetField(ctx, dto.FuncKey, slog.StringValue("Subscribe"))
	oplog := httplog.LogEntry(ctx)

	query, args, err := squirrel.
		Insert(notificationTable).
		Columns(notificationsFields...).
		Values(source, id).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		oplog.Debug("Failed to build query", "err", err.Error())
		return apperrors.ErrCouldNotBuildQuery
	}
	oplog.Debug("Query built")

	_, err = s.db.Exec(query, args...)
	if err != nil {
		oplog.Debug("Notification subscription insert failed", "err", err.Error())
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				oplog.Debug("Returning conflict")
				return apperrors.ErrSubscriptionAlreadyExists
			}
		}
		return apperrors.ErrSubscriptionNotCreated
	}
	oplog.Debug("Subscribed to birthday")

	return nil
}

func (s *PgUserStorage) Unsubscribe(ctx context.Context, source uint64, id uint64) error {
	httplog.LogEntrySetField(ctx, dto.StepKey, slog.StringValue(step))
	httplog.LogEntrySetField(ctx, dto.FuncKey, slog.StringValue("Unsubscribe"))
	oplog := httplog.LogEntry(ctx)

	query, args, err := squirrel.
		Delete(notificationTable).
		Where(squirrel.And{
			squirrel.Eq{idSourceField: source},
			squirrel.Eq{idSubscriberField: id},
		}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		oplog.Debug("Failed to build query", "err", err.Error())
		return apperrors.ErrCouldNotBuildQuery
	}
	oplog.Debug("Query built")

	_, err = s.db.Exec(query, args...)
	if err != nil {
		oplog.Debug("Notification subscription delete failed", "err", err.Error())
		return apperrors.ErrSubscriptionNotDeleted
	}
	oplog.Debug("Unsubscribed from birthday")

	return nil
}
