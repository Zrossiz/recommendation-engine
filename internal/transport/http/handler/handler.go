package handler

import (
	"engine/internal/service"

	"go.uber.org/zap"
)

type Handler struct {
	Category         *CategoryHandler
	User             *UserHandler
	Content          *ContentHandler
	UserInteractions *UserInteractionsHandler
}

type Service struct {
	UserService             UserService
	CategoryService         CategoryService
	ContentService          ContentService
	UserInteractionsService UserInteractionsService
}

func New(
	serv *service.Service,
	log *zap.Logger,
) *Handler {
	return &Handler{
		Category:         NewCategoryHandler(serv.Category, log),
		User:             NewUserHandler(serv.User, log),
		Content:          NewContentHandler(serv.Content, log),
		UserInteractions: NewUserInteractionsHandler(serv.UserInteractions, log),
	}
}
