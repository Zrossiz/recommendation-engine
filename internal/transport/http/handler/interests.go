package handler

import "go.uber.org/zap"

type InterestsHandler struct {
	service InterestsService
	log     *zap.Logger
}

type InterestsService interface {
}

func NewInterestsHandler(service InterestsService, log *zap.Logger) *InterestsHandler {
	return &InterestsHandler{
		service: service,
		log:     log,
	}
}
