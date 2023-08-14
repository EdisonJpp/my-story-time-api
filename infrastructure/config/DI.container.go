package config

import (
	"go.uber.org/fx"
	"mytimes-api/infrastructure/config/db"
)

var DIContainer = fx.Options(
	fx.Provide(db.ProvideMongoClient),
	fx.Provide(db.ProvideDatabases),
	fx.Provide(db.ProvideCollections),
)
