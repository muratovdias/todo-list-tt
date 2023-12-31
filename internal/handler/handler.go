package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/muratovdias/todo-list-tt/internal/config"
	"github.com/muratovdias/todo-list-tt/internal/service"
	"golang.org/x/exp/slog"
)

type Handler struct {
	service   service.Service
	logger    *slog.Logger
	validator *validator.Validate
}

func NewHandler(service service.Service, logger *slog.Logger) *Handler {
	return &Handler{
		service:   service,
		logger:    logger,
		validator: validator.New(),
	}
}

func (h *Handler) Routes(cfg *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	})

	app.Post("/api/todo-list/tasks", h.CreateTask)
	app.Put("/api/todo-list/tasks/:id", h.UpdateTask)
	app.Delete("/api/todo-list/tasks/:id", h.DeleteTask)
	app.Put("/api/todo-list/tasks/:id/done", h.MakeTaskDone)
	app.Get("/api/todo-list/tasks", h.TaskList)

	return app
}
