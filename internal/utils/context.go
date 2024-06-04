package utils

import (
	"birthdays/internal/apperrors"
	"birthdays/internal/pkg/dto"
	"context"
	"log/slog"
)

func GetUser(ctx context.Context) (uint64, string, error) {
	if ctx == nil {
		return 0, "", apperrors.ErrNilContext
	}

	var id uint64
	var username string
	var ok bool

	if id, ok = ctx.Value(dto.CtxUserIDKey).(uint64); !ok {
		return id, username, apperrors.ErrUserIDMissing
	}
	if username, ok = ctx.Value(dto.CtxUserKey).(string); !ok {
		return id, username, apperrors.ErrUsernameMissing
	}

	return id, username, nil
}

func GetIDParam(ctx context.Context) (uint64, error) {
	if ctx == nil {
		return 0, apperrors.ErrNilContext
	}

	var id uint64
	var ok bool

	if id, ok = ctx.Value(dto.CtxIDParamKey).(uint64); !ok {
		return id, apperrors.ErrParamMissing
	}

	return id, nil
}

func GetLogger(ctx context.Context) (*slog.Logger, error) {
	if ctx == nil {
		return nil, apperrors.ErrNilContext
	}

	logger, ok := ctx.Value(dto.CtxLoggerKey).(*slog.Logger)
	if !ok {
		return nil, apperrors.ErrLoggerMissing
	}

	return logger, nil
}
