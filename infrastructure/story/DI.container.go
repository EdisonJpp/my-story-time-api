package story

import "go.uber.org/fx"

var DIContainer = fx.Options(
	fx.Provide(NewStoryRepository),
)
