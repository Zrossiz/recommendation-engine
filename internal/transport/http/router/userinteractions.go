package router

type UserInteractionsHandler interface {
}

type UserInteractionsRouter struct {
	handler UserInteractionsHandler
}

func NewUserUnteractionsRouter(h UserInteractionsHandler) *UserInteractionsRouter {
	return &UserInteractionsRouter{
		handler: h,
	}
}
