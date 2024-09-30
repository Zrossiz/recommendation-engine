package router

import (
	"engine/internal/transport/http/handler"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct{}

func NewUserRouter() *UserRouter {
	return &UserRouter{}
}

func (u *UserRouter) RegisterRoutes(r chi.Router, h *handler.UserHandler) {
	r.Route("/user", func(r chi.Router) {

	})
}
