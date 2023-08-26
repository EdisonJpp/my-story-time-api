package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"my-story-time-api/internal/domain/shared"
)

var ErrorHandlerMiddleware = func(c *fiber.Ctx, err error) error {
	if err != nil {
		var customErr *shared.HttpException

		if errors.As(err, &customErr) {
			return c.Status(customErr.Code).JSON(fiber.Map{
				"code":    customErr.Code,
				"error":   customErr.Error(),
				"message": customErr.Message,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"error":   err.Error(),
			"message": "INTERNAL_SERVER_ERROR",
		})
	}

	return nil
}
