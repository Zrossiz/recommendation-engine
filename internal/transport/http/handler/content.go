package handler

import "engine/internal/service"

type ContentHandler struct {
}

func NewContentHandler(serv *service.ContentService) *ContentHandler {
	return &ContentHandler{}
}
