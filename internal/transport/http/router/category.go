package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CategoryRouter struct {
	handler CategoryHandler
}

type CategoryHandler interface {
	Create(rw http.ResponseWriter, r *http.Request)
}

func NewCategoryRouter(h CategoryHandler) *CategoryRouter {
	return &CategoryRouter{
		handler: h,
	}
}

func (c *CategoryRouter) RegisterRoutes(r chi.Router) {
	r.Route("/category", func(r chi.Router) {
		r.Post("/", c.handler.Create)
	})
}
