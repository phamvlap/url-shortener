package errors

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"-"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

func NotFoundError(message string, err error) *AppError {
	return &AppError{
		Code:    "not_found",
		Message: message,
		Status:  http.StatusNotFound,
		Err:     err,
	}
}

func BadRequestError(message string, err error) *AppError {
	return &AppError{
		Code:    "bad_request",
		Message: message,
		Status:  http.StatusBadRequest,
		Err:     err,
	}
}

func UnauthorizedError(message string, err error) *AppError {
	return &AppError{
		Code:    "unauthorized",
		Message: message,
		Status:  http.StatusUnauthorized,
		Err:     err,
	}
}

func InternalServerError(message string, err error) *AppError {
	return &AppError{
		Code:    "internal_server_error",
		Message: message,
		Status:  http.StatusInternalServerError,
		Err:     err,
	}
}

func ErrorHandler(c fiber.Ctx, err error) error {
	var appErr *AppError

	if errors.As(err, &appErr) {
		return c.Status(appErr.Status).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    appErr.Code,
				"message": appErr.Message,
			},
		})
	}

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		return c.Status(fiberErr.Code).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "fiber_error",
				"message": fiberErr.Message,
			},
		})
	}

	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
		"error": fiber.Map{
			"code":    "internal_server_error",
			"message": "An unexpected error occurred",
		},
	})
}
