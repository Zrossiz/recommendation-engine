package handler

import "go.uber.org/zap"

type UserInteractionsHandler struct {
	service UserInteractionsService
	log     *zap.Logger
}

type UserInteractionsService interface {
}

func NewUserInteractionsHandler(serv UserInteractionsService, log *zap.Logger) *UserInteractionsHandler {
	return &UserInteractionsHandler{
		log:     log,
		service: serv,
	}
}
