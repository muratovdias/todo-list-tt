package handler

import "github.com/muratovdias/todo-list-tt/internal/service"

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}
