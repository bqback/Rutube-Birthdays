package handlers

import (
	"birthdays/internal/apperrors"
	"birthdays/internal/pkg/dto"
	"birthdays/internal/services"
	"encoding/json"
	"log/slog"

	"net/http"

	"github.com/go-chi/httplog/v2"
)

type AuthHandler struct {
	as services.IAuthService
	us services.IUserService
}

// @Summary Авторизоваться
// @Description
// @Tags Авторизация
//
// @Accept  json
// @Produce  json
//
// @Param loginInfo body dto.LoginInfo true "Данные для авторизации"
//
// @Header 200  {string}  "Bearer token"
// @Failure 400
// @Failure 401
// @Failure 500
//
// @Router /auth/login [post]
func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	httplog.LogEntrySetField(ctx, dto.StepKey, slog.StringValue(step))
	httplog.LogEntrySetField(ctx, dto.FuncKey, slog.StringValue("Auth"))
	oplog := httplog.LogEntry(ctx)

	var info dto.LoginInfo
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		oplog.Error("Failed to decode request", "err", err.Error())
		RespondError(http.StatusBadRequest, w, ctx)
	}
	oplog.Debug("Request decoded")

	token, err := ah.as.Auth(ctx, info)
	switch err {
	case apperrors.ErrEmptyResult:
		RespondError(http.StatusUnauthorized, w, ctx)
	case apperrors.ErrWrongPassword:
		RespondError(http.StatusUnauthorized, w, ctx)
	case nil:
		w.Header().Set("Authorization", "Bearer "+token.Token)
		w.WriteHeader(http.StatusNoContent)
		oplog.Info("User authorized", dto.UserKey, token.User)
	default:
		RespondError(http.StatusInternalServerError, w, ctx)
	}

	r.Body.Close()
}

// @Summary Зарегистрироваться
// @Description
// @Tags Авторизация
//
// @Accept  json
// @Produce  json
//
// @Param signupInfo body dto.SignupInfo true "Данные для регистрации"
//
// @Header 200  {string}  "Bearer token"
// @Success 200 {object} entities.User
// @Failure 400
// @Failure 401
// @Failure 500
//
// @Router /auth/signup [post]
func (ah *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	httplog.LogEntrySetField(ctx, dto.StepKey, slog.StringValue(step))
	httplog.LogEntrySetField(ctx, dto.FuncKey, slog.StringValue("Signup"))
	oplog := httplog.LogEntry(ctx)

	var info dto.SignupInfo
	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		oplog.Error("Failed to decode request", "err", err.Error())
		RespondError(http.StatusBadRequest, w, ctx)
	}
	oplog.Debug("Request decoded")

	response, err := ah.as.Register(ctx, info)
	switch err {
	case apperrors.ErrUsernameTaken:
		RespondError(http.StatusConflict, w, ctx)
	case nil:
		w.Header().Set("Authorization", "Bearer "+response.Token)
		json.NewEncoder(w).Encode(response.User)
		w.WriteHeader(http.StatusOK)
		oplog.Info("User created", dto.UserKey, response.User)
	default:
		RespondError(http.StatusInternalServerError, w, ctx)
	}

	r.Body.Close()

}

func (ah *AuthHandler) GetAuthService() services.IAuthService {
	return ah.as
}
