package handler

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/muratovdias/todo-list-tt/internal/models"
	"github.com/muratovdias/todo-list-tt/internal/service"
	"strings"
)

func (h *Handler) CreateTask(ctx *fiber.Ctx) error {
	var todo models.ToDo

	if err := ctx.BodyParser(&todo); err != nil {
		h.logger.Error("parsing todo json", err.Error())
		return sendResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	todo.Title = strings.TrimSpace(todo.Title)
	todo.ActiveAt = strings.TrimSpace(todo.ActiveAt)

	if err := h.validator.Struct(todo); err != nil {
		h.logger.Error(err.Error())
		return sendResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	id, err := h.service.CreateTask(todo)

	if err != nil {
		h.logger.Error(err.Error())

		if errors.Is(err, service.ErrInvalidDate) || errors.Is(err, service.ErrInvalidTitle) {
			return sendResponse(ctx, fiber.StatusBadRequest, errors.Unwrap(err).Error())
		}

		return sendResponse(ctx, fiber.StatusInternalServerError, errors.Unwrap(err).Error())
	}

	h.logger.Info(fmt.Sprintf("created document id %s", id))
	return sendResponse(ctx, fiber.StatusCreated, fmt.Sprintf("created, id: %s", id))
}
