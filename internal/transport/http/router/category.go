package router

import (
	"github.com/go-chi/chi/v5"
)

type CategoryRouter struct{}

func NewCategoryRouter() *CategoryRouter {
	return &CategoryRouter{}
}

func (c *CategoryRouter) RegisterRoutes(r chi.Router) {
	r.Route("/category", func(r chi.Router) {

	})
}
