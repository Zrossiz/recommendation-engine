package router

import (
	"engine/internal/transport/http/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	Category *CategoryRouter
	User     *UserRouter
	Content  *ContentRouter
}

func New(h *handler.Handler) http.Handler {
	r := chi.NewRouter()

	router := &Router{
		Category: NewCategoryRouter(),
		User:     NewUserRouter(),
		Content:  NewContentRouter(),
	}

	router.Category.RegisterRoutes(r, h.Category)
	router.User.RegisterRoutes(r, h.User)
	router.Content.RegisterRoutes(r, h.Content)

	return r
}
