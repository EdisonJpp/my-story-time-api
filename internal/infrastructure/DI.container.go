package infrastructure

import (
	"my-story-time-api/internal/infrastructure/config"
	email_sender "my-story-time-api/internal/infrastructure/email-sender"
	"my-story-time-api/internal/infrastructure/storage"
	"my-story-time-api/internal/infrastructure/story"

	"go.uber.org/fx"
)

var DIContainer = fx.Options(
	config.DIContainer,
	story.DIContainer,
	storage.DIContainer,
	email_sender.DIContainer,
)
