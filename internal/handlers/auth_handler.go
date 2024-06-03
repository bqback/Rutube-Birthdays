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
// @Success 200  {object}  dto.JWT "JWT-токен"
// @Failure 400  {object}  apperrors.ErrorResponse
// @Failure 401  {object}  apperrors.ErrorResponse
// @Failure 500  {object}  apperrors.ErrorResponse
//
// @Router /auth/ [post]
func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	httplog.LogEntrySetFields(ctx, map[string]interface{}{
		dto.StepKey: slog.StringValue(step),
		dto.FuncKey: slog.StringValue("Auth"),
	})
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
		w.WriteHeader(http.StatusOK)
		oplog.Info("User authorized", dto.UserKey, token.User)
	default:
		RespondError(http.StatusInternalServerError, w, ctx)
	}

	r.Body.Close()
}

func (ah *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {

}

func (ah *AuthHandler) GetAuthService() services.IAuthService {
	return ah.as
}
