package application

import (
	"my-story-time-api/application/auth"
	"my-story-time-api/application/story"

	"go.uber.org/fx"
)

var DIContainer = fx.Options(
	auth.DIContainer,
	story.DIContainer,
)
