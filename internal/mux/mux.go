package mux

import (
	"birthdays/internal/handlers"
	mw "birthdays/internal/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetUpMux(manager *handlers.Handlers, logger *httplog.Logger) (http.Handler, error) {
	mux := chi.NewRouter()

	mux.Use(httplog.RequestLogger(logger))

	mux.Route("/", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", manager.AuthHandler.Login)
			r.Post("/signup", manager.AuthHandler.Signup)
		})
		r.Route("/user", func(r chi.Router) {
			r.Use(mw.AuthJWTMiddleware(manager.AuthHandler.GetAuthService()))
			r.Get("/", manager.UserHandler.GetAll)
			r.Route("/{id}", func(r chi.Router) {
				r.Post("/subscribe", manager.UserHandler.Subscribe)
				r.Post("/unsubscribe", manager.UserHandler.Unsubscribe)
			})
		})
	})

	mux.Route("/swagger/", func(r chi.Router) {
		r.Get("/*", httpSwagger.Handler(
			httpSwagger.URL("swagger/doc.json"),
		))
	})

	return mux, nil
}
