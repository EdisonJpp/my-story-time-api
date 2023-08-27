package controllers

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"mime/multipart"
	use_cases2 "my-story-time-api/internal/application/story/use-cases"
	"my-story-time-api/internal/domain/shared"
	"my-story-time-api/internal/domain/story"
	storyDto "my-story-time-api/internal/presentation/http/dto"
	"my-story-time-api/internal/presentation/http/middleware"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type StoryController struct {
	app                *fiber.App
	getStoriesUseCase  use_cases2.GetStoriesUseCase
	getStoryUseCase    use_cases2.GetStoryUseCase
	createStoryUseCase use_cases2.CreateStoryUseCase
	validator          *validator.Validate
}

func NewStoryController(
	app *fiber.App,
	getStoriesUseCase use_cases2.GetStoriesUseCase,
	getStoryUseCase use_cases2.GetStoryUseCase,
	createStoryUseCase use_cases2.CreateStoryUseCase,
	validator *validator.Validate,
) *StoryController {
	storyController := StoryController{
		app,
		getStoriesUseCase,
		getStoryUseCase,
		createStoryUseCase,
		validator,
	}

	registerStoryRoutes(&storyController)

	return &storyController
}

func registerStoryRoutes(sc *StoryController) {
	sc.app.Get(
		"/stories/:id",
		middleware.InputValidatorMiddleware(sc.validator, middleware.ParamsValidator, &storyDto.GetStoryRequestDto{}),
		sc.getStory,
	)

	sc.app.Get("/stories", sc.getStories)
	sc.app.Post("/stories", sc.createStory)
}

func (c *StoryController) getStory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	item, err := c.getStoryUseCase.Execute(id)

	if err != nil {
		if errors.Is(err, story.ErrStoryNotFound) {
			return &shared.HttpException{Code: fiber.StatusNotFound, Message: err.Error()}
		}

		return err
	}

	return ctx.JSON(item)
}

func (c *StoryController) getStories(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	stories, _ := c.getStoriesUseCase.Execute("222", page, pageSize)

	return ctx.JSON(stories)
}

func (c *StoryController) createStory(ctx *fiber.Ctx) error {
	var file multipart.FileHeader

	userId := "one-user-id"
	caption := ctx.FormValue("caption")
	text := ctx.FormValue("description")
	isAudio := ctx.FormValue("isAudio")

	if isAudio == "1" {
		fileValue, _ := ctx.FormFile("file")

		file = *fileValue
	}

	execute, err := c.createStoryUseCase.Execute(&use_cases2.CreateStoryUseCaseRequest{
		Caption: caption,
		IsAudio: isAudio,
		Text:    text,
		UserId:  userId,
		File:    &file,
	})

	if err != nil {
		return ctx.SendString(err.Error())
	}

	return ctx.JSON(execute)
}
