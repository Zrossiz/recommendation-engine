package handler

import (
	"engine/internal/service"
	"net/http"
)

type ContentHandler struct {
}

func NewContentHandler(serv *service.ContentService) *ContentHandler {
	return &ContentHandler{}
}

func (co *ContentHandler) Create(rw http.ResponseWriter, r *http.Request) {

}
