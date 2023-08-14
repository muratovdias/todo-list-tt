package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muratovdias/todo-list-tt/internal/config"
	"github.com/muratovdias/todo-list-tt/internal/service"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Routes(cfg *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	})

	return app
}
