package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

type DB struct {
	myTimesDB *mongo.Database
}

func ProvideMongoClient(lc fx.Lifecycle) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://edisonjpp:ParWWP2It24jWOUY@cluster0.1tvig.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			return client.Disconnect(context.Background())
		},
	})

	return client, nil
}
