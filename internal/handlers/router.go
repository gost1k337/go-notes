package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	_ "tg_bot/docs"
	"tg_bot/internal/service"
	"tg_bot/pkg/lib"
)

func NewRouter(services *service.Services, logger lib.Logger) *chi.Mux {
	corsCfg := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods: []string{
			http.MethodOptions,
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders: []string{"Accept", "Content-Type", "Accept-Encoding"},
	})

	r := chi.NewRouter()

	r.Use(corsCfg.Handler)
	r.Use(middleware.DefaultLogger)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"),
	))

	r.Route("/api", func(r chi.Router) {
		NewNoteRoutes(r, services, logger)
	})

	return r
}
