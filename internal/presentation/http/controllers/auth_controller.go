package controllers

import (
	authUseCases "my-story-time-api/internal/application/auth/use_cases"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	app           *fiber.App
	signInUseCase authUseCases.SignInUseCase
}

func NewAuthController(
	app *fiber.App,
	signInUseCase authUseCases.SignInUseCase,
) *AuthController {
	authController := AuthController{app, signInUseCase}

	authController.registerAuthRoutes()

	return &authController
}

func (authController *AuthController) registerAuthRoutes() {
	authController.app.Post("/sign-in", authController.signIn())
}

func (authController *AuthController) signIn() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString(authController.signInUseCase.Execute("el pepe"))
	}
}
