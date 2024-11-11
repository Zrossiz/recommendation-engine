package handler

import (
	"encoding/json"
	"engine/internal/dto"
	"net/http"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type ContentHandler struct {
	service  ContentService
	log      *zap.Logger
	validate *validator.Validate
}

type ContentService interface {
	Create(contentDTO dto.Content) (bool, error)
}

func NewContentHandler(serv ContentService, log *zap.Logger) *ContentHandler {
	return &ContentHandler{
		service:  serv,
		log:      log,
		validate: validator.New(),
	}
}

func (co *ContentHandler) Create(rw http.ResponseWriter, r *http.Request) {
	var contentDTO dto.Content

	err := co.validate.Struct(contentDTO)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&contentDTO)
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
