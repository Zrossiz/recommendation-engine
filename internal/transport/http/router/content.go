package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ContentRouter struct {
	handler ContentHandler
}

type ContentHandler interface {
	Create(rw http.ResponseWriter, r *http.Request)
}

func NewContentRouter(h ContentHandler) *ContentRouter {
	return &ContentRouter{
		handler: h,
	}
}

func (c *ContentRouter) RegisterRoutes(r chi.Router) {
	r.Route("/content", func(r chi.Router) {
		r.Post("/", c.handler.Create)
	})
}
