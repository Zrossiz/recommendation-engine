package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type InterestsRouter struct {
	handler InterestsHandler
}

type InterestsHandler interface {
	Create(rw http.ResponseWriter, r *http.Request)
}

func NewIntereestsRouter(h InterestsHandler) *InterestsRouter {
	return &InterestsRouter{
		handler: h,
	}
}

func (i *InterestsRouter) RegisterRoutes(r chi.Router) {
	r.Route("/interests", func(r chi.Router) {
		r.Post("/", i.handler.Create)
	})
}
