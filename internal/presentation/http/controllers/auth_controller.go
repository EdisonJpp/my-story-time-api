package controllers

import (
	"errors"
	"github.com/go-playground/validator/v10"
	authUseCases "my-story-time-api/internal/application/auth/use_cases"
	"my-story-time-api/internal/domain/shared"
	"my-story-time-api/internal/domain/user"
	"my-story-time-api/internal/presentation/http/dto"
	"my-story-time-api/internal/presentation/http/middleware"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	app           *fiber.App
	validator     *validator.Validate
	signInUseCase authUseCases.SignInUseCase
}

func NewAuthController(
	app *fiber.App,
	validator *validator.Validate,
	signInUseCase authUseCases.SignInUseCase,
) *AuthController {
	authController := AuthController{app, validator, signInUseCase}

	authController.registerAuthRoutes()

	return &authController
}

func (authController *AuthController) registerAuthRoutes() {
	authController.app.Post(
		"/sign-in",
		middleware.InputValidatorMiddleware(
			authController.validator,
			middleware.BodyValidator,
			&dto.SignInRequestDto{},
		),
		authController.signIn,
	)
}

func (authController *AuthController) signIn(ctx *fiber.Ctx) error {
	signInDto := new(dto.SignInRequestDto)

	ctx.BodyParser(signInDto)

	response, err := authController.signInUseCase.Execute(
		signInDto.IsEmail,
		signInDto.UserName,
		signInDto.Password,
	)

	if err != nil {
		if errors.Is(err, user.ErrUserIncorrectInformation) {
			return &shared.HttpException{Code: fiber.StatusBadRequest, Message: err.Error()}
		}

		return &shared.HttpException{Code: fiber.StatusInternalServerError, Message: err.Error()}
	}

	return ctx.JSON(response)
}
