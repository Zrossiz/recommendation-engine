package router

import (
	"engine/internal/transport/http/handler"

	"github.com/go-chi/chi/v5"
)

type ContentRouter struct{}

func NewContentRouter() *ContentRouter {
	return &ContentRouter{}
}

func (c *ContentRouter) RegisterRoutes(r chi.Router, h *handler.ContentHandler) {
	r.Route("/content", func(r chi.Router) {
		r.Post("/", h.Create)
	})
}
