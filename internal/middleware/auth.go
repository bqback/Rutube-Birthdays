package middleware

import (
	"birthdays/internal/handlers"
	"birthdays/internal/pkg/dto"
	"birthdays/internal/services"
	"context"
	"log/slog"
	"strings"

	"net/http"

	"github.com/go-chi/httplog/v2"
)

func AuthJWTMiddleware(as services.IAuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			httplog.LogEntrySetField(ctx, dto.StepKey, slog.StringValue(step))
			httplog.LogEntrySetField(ctx, dto.FuncKey, slog.StringValue("Auth"))

			oplog := httplog.LogEntry(ctx)

			token := r.Header.Get("Authorization")
			if token == "" {
				handlers.RespondError(http.StatusUnauthorized, w, ctx)
				return
			}
			oplog.Debug("Got header")

			token = strings.TrimPrefix(token, "Bearer ")
			info, err := as.Validate(token)
			if err != nil {
				oplog.Error("Error validating token", "token", token, "err", err.Error())
				handlers.RespondError(http.StatusUnauthorized, w, ctx)
				return
			}
			// if info == nil {
			// 	oplog.Error("")
			// 	handlers.RespondError(http.StatusInternalServerError, w, ctx)
			// 	return
			// }

			ctx = context.WithValue(ctx, dto.CtxUserIDKey, info.ID)
			ctx = context.WithValue(ctx, dto.CtxUserKey, info.Username)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
