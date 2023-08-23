package infrastructure

import (
	"my-story-time-api/internal/infrastructure/config"
	"my-story-time-api/internal/infrastructure/storage"
	"my-story-time-api/internal/infrastructure/story"

	"go.uber.org/fx"
)

var DIContainer = fx.Options(
	config.DIContainer,
	story.DIContainer,
	storage.DIContainer,
)
