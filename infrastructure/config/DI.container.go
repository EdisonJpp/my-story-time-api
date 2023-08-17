package config

import (
	"my-story-time-api/infrastructure/config/db"

	"go.uber.org/fx"
)

var DIContainer = fx.Options(
	fx.Provide(db.NewMongoClient),
	fx.Provide(db.NewDatabases),
	fx.Provide(db.NewCollections),
)
