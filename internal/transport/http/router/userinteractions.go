package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserInteractionsHandler interface {
	Create(rw http.ResponseWriter, r *http.Request)
}

type UserInteractionsRouter struct {
	handler UserInteractionsHandler
}

func NewUserUnteractionsRouter(h UserInteractionsHandler) *UserInteractionsRouter {
	return &UserInteractionsRouter{
		handler: h,
	}
}

func (i *UserInteractionsRouter) RegisterRoutes(r chi.Router) {
	r.Route("/interaction", func(r chi.Router) {
		r.Post("/", i.handler.Create)
	})
}
