package handler

import (
	"encoding/json"
	"engine/internal/dto"
	"net/http"

	"go.uber.org/zap"
)

type ContentHandler struct {
	service ContentService
	log     *zap.Logger
}

type ContentService interface {
	Create(contentDTO dto.Content) (bool, error)
}

func NewContentHandler(serv ContentService, log *zap.Logger) *ContentHandler {
	return &ContentHandler{
		service: serv,
		log:     log,
	}
}

func (co *ContentHandler) Create(rw http.ResponseWriter, r *http.Request) {
	var contentDTO dto.Content

	err := json.NewDecoder(r.Body).Decode(&contentDTO)
	if err != nil {
		http.Error(rw, "invalid body", http.StatusBadRequest)
		return
	}

	_, err = co.service.Create(contentDTO)
	if err != nil {
		co.log.Error("create content", zap.Error(err))
		http.Error(rw, "internal server error", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}
