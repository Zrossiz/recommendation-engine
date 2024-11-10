package router

import (
	"github.com/go-chi/chi/v5"
)

type ContentRouter struct{}

func NewContentRouter() *ContentRouter {
	return &ContentRouter{}
}

func (c *ContentRouter) RegisterRoutes(r chi.Router) {
	r.Route("/content", func(r chi.Router) {
	})
}
