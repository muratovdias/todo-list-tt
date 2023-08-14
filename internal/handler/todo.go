package handler

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/muratovdias/todo-list-tt/internal/models"
	"github.com/muratovdias/todo-list-tt/internal/service"
	"github.com/muratovdias/todo-list-tt/internal/storage"
	"strings"
)

func (h *Handler) CreateTask(ctx *fiber.Ctx) error {
	var todo models.ToDo

	if err := ctx.BodyParser(&todo); err != nil {
		h.logger.Error("parsing todo json", err.Error())
		return sendResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	// удаляем все ведущие и завершающие пробелы,
	todo.Title = strings.TrimSpace(todo.Title)
	todo.ActiveAt = strings.TrimSpace(todo.ActiveAt)

	// валидируем нашу структуру
	if err := h.validator.Struct(todo); err != nil {
		h.logger.Error(err.Error())
		return sendResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	id, err := h.service.CreateTask(todo)

	if err != nil {
		h.logger.Error(err.Error())

		if errors.Is(err, service.ErrInvalidDate) || errors.Is(err, service.ErrInvalidTitle) {
			return sendResponse(ctx, fiber.StatusBadRequest, err.Error())
		}

		return sendResponse(ctx, fiber.StatusInternalServerError, errors.Unwrap(err).Error())
	}

	h.logger.Info(fmt.Sprintf("created document id %s", id))
	return sendResponse(ctx, fiber.StatusCreated, fmt.Sprintf("created, id: %s", id))
}

func (h *Handler) UpdateTask(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var todo models.ToDo
	todo.ID = id

	if err := ctx.BodyParser(&todo); err != nil {
		h.logger.Error("parsing json", err.Error())
		return sendResponse(ctx, fiber.StatusInternalServerError, err.Error())
	}

	// удаляем все ведущие и завершающие пробелы,
	todo.Title = strings.TrimSpace(todo.Title)
	todo.ActiveAt = strings.TrimSpace(todo.ActiveAt)

	// валидируем нашу структуру
	if err := h.validator.Struct(todo); err != nil {
		h.logger.Error(err.Error())
		return sendResponse(ctx, fiber.StatusBadRequest, err.Error())
	}

	res, err := h.service.UpdateTask(todo)
	if err != nil {
		h.logger.Error(err.Error())

		if errors.Is(err, service.ErrInvalidDate) || errors.Is(err, service.ErrInvalidTitle) || errors.Is(err, storage.ErrInvalidId) {
			return sendResponse(ctx, fiber.StatusBadRequest, err.Error())
		}
		return sendResponse(ctx, fiber.StatusInternalServerError, errors.Unwrap(err).Error())
	}

	h.logger.Info(fmt.Sprintf("updated documents count is %d", res))
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) DeleteTask(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	res, err := h.service.DeleteTask(id)
	if err != nil {
		h.logger.Error(err.Error())
		if errors.Is(err, storage.ErrInvalidId) {
			return sendResponse(ctx, fiber.StatusBadRequest, err.Error())
		}
		return sendResponse(ctx, fiber.StatusInternalServerError, errors.Unwrap(err).Error())
	}

	h.logger.Info(fmt.Sprintf("deleted documnets count is %d", res))
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) MakeTaskDone(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	res, err := h.service.MakeTaskDone(id)
	if err != nil {
		h.logger.Error(err.Error())

		if errors.Is(err, storage.ErrInvalidId) {
			return sendResponse(ctx, fiber.StatusBadRequest, err.Error())
		}

		return sendResponse(ctx, fiber.StatusInternalServerError, errors.Unwrap(err).Error())
	}

	h.logger.Info(fmt.Sprintf("updated status count is %d", res))
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) TaskList(ctx *fiber.Ctx) error {
	status := ctx.Query("status", "active")

	todoList, err := h.service.TaskList(status)
	if err != nil {
		h.logger.Error(err.Error())
		return sendResponse(ctx, fiber.StatusInternalServerError, errors.Unwrap(err).Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(todoList)
}
