package handler

import (
	"net/http"

	"go.uber.org/zap"
)

type ContentHandler struct {
}

type ContentService interface {
}

func NewContentHandler(serv ContentService, log *zap.Logger) *ContentHandler {
	return &ContentHandler{}
}

func (co *ContentHandler) Create(rw http.ResponseWriter, r *http.Request) {

}
