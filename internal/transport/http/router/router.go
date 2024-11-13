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
	Interests        InterestsRouter
}

func New(h *handler.Handler) http.Handler {
	r := chi.NewRouter()

	router := &Router{
		Category:         *NewCategoryRouter(h.Category),
		User:             *NewUserRouter(h.User),
		Content:          *NewContentRouter(h.Content),
		UserInteractions: *NewUserUnteractionsRouter(h.UserInteractions),
		Interests:        *NewIntereestsRouter(h.Interests),
	}

	router.Category.RegisterRoutes(r)
	router.User.RegisterRoutes(r)
	router.Content.RegisterRoutes(r)
	router.UserInteractions.RegisterRoutes(r)
	router.Interests.RegisterRoutes(r)

	return r
}
