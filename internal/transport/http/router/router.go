package router

import (
	"engine/internal/transport/http/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	Category         CategoryRouter
	User             UserRouter
	Content          ContentRouter
	UserInteractions UserInteractionsRouter
}

func New(h *handler.Handler) http.Handler {
	r := chi.NewRouter()

	router := &Router{
		Category:         *NewCategoryRouter(),
		User:             *NewUserRouter(h.User),
		Content:          *NewContentRouter(),
		UserInteractions: *NewUserUnteractionsRouter(h.UserInteractions),
	}

	router.Category.RegisterRoutes(r)
	router.User.RegisterRoutes(r)
	router.Content.RegisterRoutes(r)

	return r
}
