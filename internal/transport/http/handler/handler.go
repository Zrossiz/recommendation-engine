package handler

import "engine/internal/service"

type Handler struct {
	Category *CategoryHandler
	User     *UserHandler
	Content  *ContentHandler
}

func New(serv *service.Service) *Handler {
	return &Handler{
		Category: NewCategoryHandler(serv.Category),
		User:     NewUserHandler(serv.User),
		Content:  NewContentHandler(serv.Content),
	}
}
