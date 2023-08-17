package story

import (
	"my-story-time-api/application/story/use_cases"

	"go.uber.org/fx"
)

var DIContainer = fx.Options(
	fx.Provide(
		use_cases.NewGetStoriesUseCase,
		use_cases.NewGetStoryUseCase,
		use_cases.NewCreateStoryUseCase,
	),
)
