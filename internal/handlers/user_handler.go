package handlers

import "github.com/mamun-jsx/user-management-golang/internal/service"

type UserHandler struct {
	service *service.UserService
}

// make a new handler

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// 
