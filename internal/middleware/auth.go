package middleware

import (
	"birthdays/internal/services"

	"net/http"
)

func AuthMiddleware(as services.IAuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		})
	}
}
