package handlers

import (
	"birthdays/internal/apperrors"
	"birthdays/internal/pkg/dto"
	"birthdays/internal/services"
	"birthdays/internal/utils"
	"encoding/json"
	"log/slog"

	"net/http"

	"github.com/go-chi/httplog/v2"
)

type UserHandler struct {
	us services.IUserService
}

func (uh *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	httplog.LogEntrySetField(ctx, dto.StepKey, slog.StringValue(step))
	httplog.LogEntrySetField(ctx, dto.FuncKey, slog.StringValue("Auth"))

	id, username, err := utils.GetUserFromContext(ctx)
	switch err {
	case nil:
		break
	case apperrors.ErrNilContext:
		RespondError(http.StatusInternalServerError, w, ctx)
		r.Body.Close()
		return
	default:
		RespondError(http.StatusUnauthorized, w, ctx)
		r.Body.Close()
		return
	}

	httplog.LogEntrySetField(ctx, dto.UserIDKey, slog.Uint64Value(id))
	httplog.LogEntrySetField(ctx, dto.UserKey, slog.StringValue(username))

	oplog := httplog.LogEntry(ctx)
	oplog.Info("User list requested")

	users, err := uh.us.GetAll(ctx)
	switch err {
	case apperrors.ErrEmptyResult:
		fallthrough
	case nil:
		json.NewEncoder(w).Encode(users)
		w.WriteHeader(http.StatusOK)
		oplog.Info("User list sent")
	default:
		RespondError(http.StatusInternalServerError, w, ctx)
	}

	r.Body.Close()
}

func (uh *UserHandler) Subscribe(w http.ResponseWriter, r *http.Request) {

}

func (uh *UserHandler) Unsubscribe(w http.ResponseWriter, r *http.Request) {

}
