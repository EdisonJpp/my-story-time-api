package application

import (
	"go.uber.org/fx"
	"mytimes-api/application/auth"
	"mytimes-api/application/story"
)

var DIContainer = fx.Options(
	auth.DIContainer,
	story.DIContainer,
)
