package handler

import "go.uber.org/zap"

type CategoryHandler struct {
}

type CategoryService interface {
}

func NewCategoryHandler(serv CategoryService, log *zap.Logger) *CategoryHandler {
	return &CategoryHandler{}
}
