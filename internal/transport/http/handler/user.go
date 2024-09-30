package handler

import "engine/internal/service"

type UserHandler struct {
}

func NewUserHandler(serv *service.UserService) *UserHandler {
	return &UserHandler{}
}
