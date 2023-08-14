package handler

import "github.com/gofiber/fiber/v2"

type response struct {
	Message string `json:"message"`
}

func sendResponse(ctx *fiber.Ctx, status int, message string) error {
	return ctx.Status(status).JSON(response{Message: message})
}
