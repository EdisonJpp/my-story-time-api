package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"my-story-time-api/internal/domain/shared"
)

func ProtectedMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   "a-sign-key",
		ErrorHandler: jwtError,
	})
}

func jwtError(_ *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return &shared.HttpException{Message: "MISSING_OR_MALFORMED_TOKEN", Code: fiber.StatusBadRequest}
	}

	return &shared.HttpException{Message: "INVALID_OR_EXPIRED_TOKEN", Code: fiber.StatusBadRequest}
}
