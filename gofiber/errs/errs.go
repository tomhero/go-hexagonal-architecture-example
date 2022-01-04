package errs

import (
	"github.com/gofiber/fiber/v2"
)

func NewNotFoundError(message string) *fiber.Error {
	return fiber.NewError(fiber.StatusNotFound, message)
}

func NewUnexpectedError() *fiber.Error {
	return fiber.ErrInternalServerError
}

func NewValidationError(message string) *fiber.Error {
	return fiber.NewError(fiber.StatusUnprocessableEntity, message)
}
