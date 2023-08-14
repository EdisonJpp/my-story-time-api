package controllers

import (
	"github.com/gofiber/fiber/v2"
	authUseCases "mytimes-api/application/auth/use_cases"
)

type AuthController struct {
	app           *fiber.App
	signInUseCase authUseCases.SignInUseCase
}

func ProvideAuthController(
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
