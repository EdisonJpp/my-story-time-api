package controllers

import (
	"errors"
	"mime/multipart"
	"my-story-time-api/application/story/use_cases"
	"my-story-time-api/domain/story"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type StoryController struct {
	app                *fiber.App
	getStoriesUseCase  use_cases.GetStoriesUseCase
	getStoryUseCase    use_cases.GetStoryUseCase
	createStoryUseCase use_cases.CreateStoryUseCase
}

func NewStoryController(
	app *fiber.App,
	getStoriesUseCase use_cases.GetStoriesUseCase,
	getStoryUseCase use_cases.GetStoryUseCase,
	createStoryUseCase use_cases.CreateStoryUseCase,
) *StoryController {
	storyController := StoryController{
		app,
		getStoriesUseCase,
		getStoryUseCase,
		createStoryUseCase,
	}

	registerStoryRoutes(&storyController)

	return &storyController
}

func registerStoryRoutes(storyController *StoryController) {
	storyController.app.Get("/stories", storyController.getStories)
	storyController.app.Get("/stories/:id", storyController.getStory)
	storyController.app.Post("/stories", storyController.createStory)
}

func (c *StoryController) getStories(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	stories, _ := c.getStoriesUseCase.Execute("222", page, pageSize)

	return ctx.JSON(stories)
}

func (c *StoryController) getStory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).SendString("ID_IS_REQUIRED")
	}

	item, err := c.getStoryUseCase.Execute(id)

	if err != nil {
		if errors.Is(err, story.ErrStoryNotFound) {
			return ctx.Status(fiber.StatusNotFound).SendString(err.Error())
		}

		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(item)
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

	execute, err := c.createStoryUseCase.Execute(&use_cases.CreateStoryUseCaseRequest{
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
