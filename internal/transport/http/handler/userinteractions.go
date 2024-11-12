package handler

import (
	"encoding/json"
	"engine/internal/dto"
	"net/http"

	"go.uber.org/zap"
)

type UserInteractionsHandler struct {
	service UserInteractionsService
	log     *zap.Logger
}

type UserInteractionsService interface {
	Create(dto dto.CreateInteraction) (bool, error)
}

func NewUserInteractionsHandler(serv UserInteractionsService, log *zap.Logger) *UserInteractionsHandler {
	return &UserInteractionsHandler{
		log:     log,
		service: serv,
	}
}

func (i *UserInteractionsHandler) Create(rw http.ResponseWriter, r *http.Request) {
	var interactionDTO dto.CreateInteraction

	err := json.NewDecoder(r.Body).Decode(&interactionDTO)
	if err != nil {
		http.Error(rw, "invalid body", http.StatusBadRequest)
		return
	}

	_, err = i.service.Create(interactionDTO)
	if err != nil {
		i.log.Error("create interaction", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}
