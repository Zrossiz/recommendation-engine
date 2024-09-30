package router

import (
	"engine/internal/transport/http/handler"

	"github.com/go-chi/chi/v5"
)

type CategoryRouter struct{}

func NewCategoryRouter() *CategoryRouter {
	return &CategoryRouter{}
}

func (c *CategoryRouter) RegisterRoutes(r chi.Router, h *handler.CategoryHandler) {
	r.Route("/category", func(r chi.Router) {

	})
}
