package infrastructure

import (
	"my-story-time-api/infrastructure/config"
	"my-story-time-api/infrastructure/storage"
	"my-story-time-api/infrastructure/story"

	"go.uber.org/fx"
)

var DIContainer = fx.Options(
	config.DIContainer,
	story.DIContainer,
	storage.DIContainer,
)
