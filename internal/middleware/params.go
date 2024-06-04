package middleware

import (
	"birthdays/internal/handlers"
	"birthdays/internal/pkg/dto"
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
)

func IDFromUrl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		httplog.LogEntrySetField(ctx, dto.StepKey, slog.StringValue(step))
		httplog.LogEntrySetField(ctx, dto.FuncKey, slog.StringValue("IDFromUrl"))

		oplog := httplog.LogEntry(ctx)

		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			oplog.Error("Failed to convert param id to uint", "err", err.Error())
			handlers.RespondError(http.StatusBadRequest, w, ctx)
			return
		}

		rCtx := context.WithValue(r.Context(), dto.CtxIDParamKey, id)
		oplog.Debug("Stored id in context", "id", id)

		next.ServeHTTP(w, r.WithContext(rCtx))
	})
}
