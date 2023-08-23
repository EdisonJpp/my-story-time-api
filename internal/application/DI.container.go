package application

import (
	"my-story-time-api/internal/application/auth"
	"my-story-time-api/internal/application/story"

	"go.uber.org/fx"
)

var DIContainer = fx.Options(
	auth.DIContainer,
	story.DIContainer,
)
