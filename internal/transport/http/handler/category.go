package handler

import "engine/internal/service"

type CategoryHandler struct {
}

func NewCategoryHandler(serv *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{}
}
