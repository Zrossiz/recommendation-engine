package handler

import (
	"encoding/json"
	"engine/internal/dto"
	"net/http"

	"go.uber.org/zap"
)

type InterestsHandler struct {
	service InterestsService
	log     *zap.Logger
}

type InterestsService interface {
	Create(interestDTO dto.Interest) (bool, error)
}

func NewInterestsHandler(service InterestsService, log *zap.Logger) *InterestsHandler {
	return &InterestsHandler{
		service: service,
		log:     log,
	}
}

func (i *InterestsHandler) Create(rw http.ResponseWriter, r *http.Request) {
	var interestDTO dto.Interest

	err := json.NewDecoder(r.Body).Decode(&interestDTO)
	if err != nil {
		http.Error(rw, "invalid body", http.StatusBadRequest)
		return
	}

	_, err = i.service.Create(interestDTO)
	if err != nil {
		i.log.Error("create interest", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}
