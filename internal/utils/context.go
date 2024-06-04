package utils

import (
	"birthdays/internal/apperrors"
	"birthdays/internal/pkg/dto"
	"context"
)

func GetUserFromContext(ctx context.Context) (uint64, string, error) {
	if ctx == nil {
		return 0, "", apperrors.ErrNilContext
	}

	var id uint64
	var username string
	var ok bool

	if id, ok = ctx.Value(dto.UserIDKey).(uint64); !ok {
		return id, username, apperrors.ErrUserIDMissing
	}
	if username, ok = ctx.Value(dto.UserKey).(string); !ok {
		return id, username, apperrors.ErrUsernameMissing
	}

	return id, username, nil
}
