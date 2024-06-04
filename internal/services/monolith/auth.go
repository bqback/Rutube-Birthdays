package monolith

import (
	"birthdays/internal/apperrors"
	"birthdays/internal/auth"
	"birthdays/internal/pkg/dto"
	"birthdays/internal/pkg/entities"
	"birthdays/internal/storages"
	"birthdays/internal/utils"
	"context"
	"log/slog"

	"github.com/go-chi/httplog/v2"
)

type AuthJWTService struct {
	manager *auth.AuthManager
	storage storages.IAuthStorage
}

func NewAuthJWTService(storage storages.IAuthStorage, manager *auth.AuthManager) *AuthJWTService {
	return &AuthJWTService{
		storage: storage,
		manager: manager,
	}
}

func (s *AuthJWTService) Auth(ctx context.Context, info dto.LoginInfo) (*entities.JWT, error) {
	httplog.LogEntrySetField(ctx, dto.StepKey, slog.StringValue(step))
	oplog := httplog.LogEntry(ctx)

	user, err := s.storage.GetByUsername(ctx, info.Username)
	if err != nil {
		return nil, err
	}
	oplog.Debug("Got user")

	if err = utils.ComparePasswords(user.PasswordHash, info.Password); err != nil {
		return nil, apperrors.ErrWrongPassword
	}
	oplog.Debug("Passwords match")

	token, err := s.manager.GenerateToken(&dto.TokenInfo{ID: user.ID, Username: info.Username})
	if err != nil {
		oplog.Error("Error generating token", "err", err.Error())
		return nil, err
	}
	oplog.Debug("Token generated")

	jwt := &entities.JWT{
		Token: token,
		User:  user.Username,
	}

	return jwt, nil
}

func (s *AuthJWTService) Register(ctx context.Context, info dto.SignupInfo) (*dto.SignupResponse, error) {
	httplog.LogEntrySetField(ctx, dto.StepKey, slog.StringValue(step))
	httplog.LogEntrySetField(ctx, dto.FuncKey, slog.StringValue("Register"))
	oplog := httplog.LogEntry(ctx)

	_, err := s.storage.GetByUsername(ctx, info.Username)
	if err == nil {
		return nil, apperrors.ErrUsernameTaken
	}
	oplog.Debug("User doesn't exist")

	info.Password, err = utils.HashFromPassword(info.Password)
	if err != nil {
		oplog.Error("Error hashing password", "err", err.Error())
		return nil, err
	}

	user, err := s.storage.Create(ctx, info)
	if err != nil {
		return nil, err
	}
	oplog.Debug("Signed up user")

	token, err := s.manager.GenerateToken(&dto.TokenInfo{ID: user.ID, Username: info.Username})
	if err != nil {
		oplog.Error("Error generating token", "err", err.Error())
		return nil, err
	}
	oplog.Debug("Token generated")

	response := &dto.SignupResponse{
		Token: token,
		User:  *user,
	}

	return response, nil
}

func (s *AuthJWTService) Validate(token string) (*dto.TokenInfo, error) {
	return s.manager.ValidateToken(token)
}
