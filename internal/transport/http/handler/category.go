package handler

import (
	"encoding/json"
	"engine/internal/apperrors"
	"engine/internal/dto"
	"net/http"

	"go.uber.org/zap"
)

type CategoryHandler struct {
	service CategoryService
	log     *zap.Logger
}

type CategoryService interface {
	Create(dto dto.Category) (bool, error)
}

func NewCategoryHandler(serv CategoryService, log *zap.Logger) *CategoryHandler {
	return &CategoryHandler{
		service: serv,
		log:     log,
	}
}

func (c *CategoryHandler) Create(rw http.ResponseWriter, r *http.Request) {
	var categoryDTO dto.Category

	err := json.NewDecoder(r.Body).Decode(&categoryDTO)
	if err != nil {
		http.Error(rw, "invalid body", http.StatusBadRequest)
		return
	}

	if categoryDTO.Name == "" {
		http.Error(rw, "name not prvoded", http.StatusBadRequest)
		return
	}

	_, err = c.service.Create(categoryDTO)
	if err != nil {
		if err == apperrors.ErrAlreadyExist {
			http.Error(rw, "category already exist", http.StatusConflict)
			return
		}

		c.log.Error("create category", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("content-type", "application/json")
}
