package infrastructure

import (
	"go.uber.org/fx"
	"mytimes-api/infrastructure/config"
	"mytimes-api/infrastructure/storage"
	"mytimes-api/infrastructure/story"
)

var DIContainer = fx.Options(
	config.DIContainer,
	story.DIContainer,
	storage.DIContainer,
)
