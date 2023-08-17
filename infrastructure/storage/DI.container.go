package storage

import "go.uber.org/fx"

var DIContainer = fx.Options(
	fx.Provide(NewStorageClient, NewStorageRepository),
)
