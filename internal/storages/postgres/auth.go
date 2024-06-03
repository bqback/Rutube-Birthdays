package postgresql

import (
	"birthdays/internal/apperrors"
	"birthdays/internal/pkg/dto"
	"birthdays/internal/pkg/entities"
	"birthdays/internal/utils"
	"context"
	"database/sql"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/go-chi/httplog/v2"
	"github.com/jmoiron/sqlx"
)

type PgAuthStorage struct {
	db *sqlx.DB
}

func NewAuthStorage(db *sqlx.DB) *PgAuthStorage {
	return &PgAuthStorage{
		db: db,
	}
}

func (s *PgAuthStorage) Auth(ctx context.Context, info dto.LoginInfo) (*dto.DBUser, error) {
	httplog.LogEntrySetFields(ctx, map[string]interface{}{
		dto.StepKey: slog.StringValue(step),
		dto.FuncKey: slog.StringValue("Auth"),
	})
	oplog := httplog.LogEntry(ctx)

	query, args, err := squirrel.
		Select(authLoginSelectFields...).
		From(authTable).
		Where(squirrel.Eq{authLoginField: info.Username}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		oplog.Debug("Failed to build query", "err", err.Error())
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	oplog.Debug("Query built")

	var user dto.DBUser
	if err := s.db.Get(&user, query, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrEmptyResult
		}
		oplog.Debug("User select failed", "err", err.Error())
		return nil, apperrors.ErrUserNotSelected
	}
	oplog.Debug("User selected")

	err = utils.ComparePasswords(user.PasswordHash, info.Password)
	if err != nil {
		oplog.Debug("Passwords don't match", "err", err.Error())
		return nil, apperrors.ErrWrongPassword
	}
	oplog.Debug("Passwords match")

	return &user, nil
}

func (s *PgAuthStorage) Register(ctx context.Context, info dto.SignupInfo) (*entities.User, error) {
	httplog.LogEntrySetFields(ctx, map[string]interface{}{
		dto.StepKey: slog.StringValue(step),
		dto.FuncKey: slog.StringValue("Register"),
	})
	oplog := httplog.LogEntry(ctx)

	tx, err := s.db.Begin()
	if err != nil {
		oplog.Debug("Failed to start transaction", "err", err.Error())
		return nil, apperrors.ErrCouldNotBeginTransaction
	}
	oplog.Debug("Transaction started")

	query1, args, err := squirrel.
		Insert(authTable).
		Columns(authSignupInsertFields...).
		Values(info.Username, info.Password).
		PlaceholderFormat(squirrel.Dollar).
		Suffix(returningIDSuffix).
		ToSql()
	if err != nil {
		oplog.Debug("Failed to build query", "err", err.Error())
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	oplog.Debug("Query built")

	var userID int
	row := tx.QueryRow(query1, args...)
	if err := row.Scan(&userID); err != nil {
		oplog.Debug("Auth insert failed", "err", err.Error())
		err = tx.Rollback()
		if err != nil {
			oplog.Error("Rollback failed", "err", err.Error())
			return nil, apperrors.ErrCouldNotRollback
		}
		return nil, apperrors.ErrUserNotCreated
	}
	oplog.Debug("Auth entry created")

	query2, args, err := squirrel.
		Insert(userTable).
		Columns(newUserInsertFields...).
		Values(userID, info.Email, info.DOB).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		oplog.Debug("Failed to build query", "err", err.Error())
		return nil, apperrors.ErrCouldNotBuildQuery
	}
	oplog.Debug("Query built")

	_, err = tx.Exec(query2, args...)
	if err != nil {
		oplog.Debug("User insert failed", "err", err.Error())
		err = tx.Rollback()
		if err != nil {
			oplog.Error("Rollback failed", "err", err.Error())
			return nil, apperrors.ErrCouldNotRollback
		}
		return nil, apperrors.ErrUserNotCreated
	}
	oplog.Debug("User entry created")

	err = tx.Commit()
	if err != nil {
		oplog.Debug("Failed to commit changes", "err", err.Error())
		err = tx.Rollback()
		if err != nil {
			oplog.Error("Rollback failed", "err", err.Error())
			return nil, apperrors.ErrCouldNotRollback
		}
		return nil, apperrors.ErrUserNotCreated
	}
	oplog.Debug("Changes commited")

	newUser := &entities.User{
		ID:      uint64(userID),
		Name:    "",
		Surname: "",
		Email:   info.Email,
		DOB:     info.DOB,
	}

	return newUser, nil
}
