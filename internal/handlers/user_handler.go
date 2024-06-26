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

// @Summary Получить список пользователей
// @Description
// @Tags Пользователи
//
// @Produce  json
//
// @Security JWT
//
// @Success 200 {object} []entities.User
// @Failure 400
// @Failure 401
// @Failure 500
//
// @Router /user [get]
func (uh *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	httplog.LogEntrySetField(ctx, dto.StepKey, slog.StringValue(step))
	httplog.LogEntrySetField(ctx, dto.FuncKey, slog.StringValue("Auth"))

	id, username, err := utils.GetUser(ctx)
	switch err {
	case nil:
		break
	case apperrors.ErrNilContext:
		httplog.LogEntry(ctx).Error("nil context")
		RespondError(http.StatusInternalServerError, w, ctx)
		r.Body.Close()
		return
	default:
		httplog.LogEntry(ctx).Error("User unauthorized")
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
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			RespondError(http.StatusInternalServerError, w, ctx)
		}
		w.WriteHeader(http.StatusOK)
		oplog.Info("User list sent")
	default:
		RespondError(http.StatusInternalServerError, w, ctx)
	}

	r.Body.Close()
}

// @Summary Подписаться на день рождения пользователя
// @Description
// @Tags Пользователи
//
// @Param id path uint true "ID пользователя"
//
// @Security JWT
//
// @Success 200
// @Failure 400
// @Failure 409
// @Failure 500
//
// @Router /user/{id}/subscribe [post]
func (uh *UserHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	httplog.LogEntrySetField(ctx, dto.StepKey, slog.StringValue(step))
	httplog.LogEntrySetField(ctx, dto.FuncKey, slog.StringValue("Subscribe"))

	id, username, err := utils.GetUser(ctx)
	switch err {
	case nil:
		break
	case apperrors.ErrNilContext:
		httplog.LogEntry(ctx).Error("nil context")
		RespondError(http.StatusInternalServerError, w, ctx)
		r.Body.Close()
		return
	default:
		httplog.LogEntry(ctx).Error("User unauthorized")
		RespondError(http.StatusUnauthorized, w, ctx)
		r.Body.Close()
		return
	}

	httplog.LogEntrySetField(ctx, dto.UserIDKey, slog.Uint64Value(id))
	httplog.LogEntrySetField(ctx, dto.UserKey, slog.StringValue(username))

	source, err := utils.GetIDParam(ctx)
	if err != nil {
		httplog.LogEntry(ctx).Error("failed to get id to subscribe to from context", "err", err.Error())
		RespondError(http.StatusInternalServerError, w, ctx)
		r.Body.Close()
		return
	}

	httplog.LogEntrySetField(ctx, dto.UserIDKey, slog.Uint64Value(source))

	oplog := httplog.LogEntry(ctx)
	oplog.Info("Subscribing to source's birthday")

	if id == source {
		oplog.Error("Source id equals subscriber's id")
		RespondError(http.StatusConflict, w, ctx)
		r.Body.Close()
		return
	}

	err = uh.us.Subscribe(ctx, source, id)
	switch err {
	case nil:
		w.WriteHeader(http.StatusOK)
		oplog.Info("Subscribed")
	case apperrors.ErrSubscriptionAlreadyExists:
		RespondError(http.StatusConflict, w, ctx)
	default:
		RespondError(http.StatusInternalServerError, w, ctx)
	}

	r.Body.Close()
}

// @Summary Отписаться от дня рождения пользователя
// @Description
// @Tags Пользователи
//
// @Param id path uint true "ID пользователя"
//
// @Security JWT
//
// @Success 200
// @Failure 400
// @Failure 500
//
// @Router /user/{id}/unsubscribe [post]
func (uh *UserHandler) Unsubscribe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	httplog.LogEntrySetField(ctx, dto.StepKey, slog.StringValue(step))
	httplog.LogEntrySetField(ctx, dto.FuncKey, slog.StringValue("Subscribe"))

	id, username, err := utils.GetUser(ctx)
	switch err {
	case nil:
		break
	case apperrors.ErrNilContext:
		httplog.LogEntry(ctx).Error("nil context")
		RespondError(http.StatusInternalServerError, w, ctx)
		r.Body.Close()
		return
	default:
		httplog.LogEntry(ctx).Error("User unauthorized")
		RespondError(http.StatusUnauthorized, w, ctx)
		r.Body.Close()
		return
	}

	httplog.LogEntrySetField(ctx, dto.UserIDKey, slog.Uint64Value(id))
	httplog.LogEntrySetField(ctx, dto.UserKey, slog.StringValue(username))

	source, err := utils.GetIDParam(ctx)
	if err != nil {
		httplog.LogEntry(ctx).Error("failed to get id to unsubscribe to from context", "err", err.Error())
		RespondError(http.StatusInternalServerError, w, ctx)
		r.Body.Close()
		return
	}

	httplog.LogEntrySetField(ctx, dto.UserIDKey, slog.Uint64Value(source))

	oplog := httplog.LogEntry(ctx)
	oplog.Info("Unsubscribing from source's birthday")

	if id == source {
		oplog.Error("Source id equals subscriber's id")
		RespondError(http.StatusConflict, w, ctx)
		r.Body.Close()
		return
	}

	err = uh.us.Unsubscribe(ctx, source, id)
	switch err {
	case nil:
		w.WriteHeader(http.StatusOK)
		oplog.Info("Unsubscribed")
	default:
		RespondError(http.StatusInternalServerError, w, ctx)
	}

	r.Body.Close()
}
