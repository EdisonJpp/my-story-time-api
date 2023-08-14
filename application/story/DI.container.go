package story

import (
	"go.uber.org/fx"
	"mytimes-api/application/story/use_cases"
)

var DIContainer = fx.Options(
	fx.Provide(
		use_cases.ProvideGetStoriesUseCase,
		use_cases.ProvideGetStoryUseCase,
		use_cases.ProvideCreateStoryUseCase,
	),
)
