package handlers

import (
	chi "github.com/go-chi/chi/v5"
	"net/http"
	"tg_bot/internal/service"
	"tg_bot/pkg/lib"
)

type Handler struct {
	services *service.Services
	http     *chi.Mux
	logger   lib.Logger
}

func New(services *service.Services, logger lib.Logger) *Handler {
	h := &Handler{
		services: services,
		logger:   logger,
	}

	h.http = NewRouter(services, logger)

	return h
}

func (h *Handler) HTTP() http.Handler {
	return h.http
}
