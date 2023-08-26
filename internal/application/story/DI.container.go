package story

import (
	use_cases2 "my-story-time-api/internal/application/story/use-cases"

	"go.uber.org/fx"
)

var DIContainer = fx.Options(
	fx.Provide(
		use_cases2.NewGetStoriesUseCase,
		use_cases2.NewGetStoryUseCase,
		use_cases2.NewCreateStoryUseCase,
	),
)
