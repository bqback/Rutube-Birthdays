package handlers

import (
	"birthdays/internal/pkg/dto"
	"birthdays/internal/services"
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/httplog/v2"
)

const step string = "handler"

type Handlers struct {
	AuthHandler
	UserHandler
}

func NewHandlers(services *services.Services) *Handlers {
	return &Handlers{
		AuthHandler: *NewAuthHandler(services.Auth, services.User),
		UserHandler: *NewUserHandler(services.User),
	}
}

func NewAuthHandler(as services.IAuthService, us services.IUserService) *AuthHandler {
	return &AuthHandler{
		as: as,
		us: us,
	}
}

func NewUserHandler(us services.IUserService) *UserHandler {
	return &UserHandler{
		us: us,
	}
}

func RespondError(code int, w http.ResponseWriter, ctx context.Context) {
	httplog.LogEntrySetFields(ctx, map[string]interface{}{
		dto.FuncKey: slog.StringValue("RespondError"),
	})
	oplog := httplog.LogEntry(ctx)

	w.WriteHeader(code)
	_, err := w.Write([]byte(http.StatusText(code)))
	if err != nil {
		oplog.Error("Failed to return error to client", "err", err.Error())
	}
}
