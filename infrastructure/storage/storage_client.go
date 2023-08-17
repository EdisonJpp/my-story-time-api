package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"go.uber.org/fx"
	"google.golang.org/api/option"
)

func NewStorageClient(lc fx.Lifecycle) *storage.Client {
	ctx := context.Background()

	client, err := storage.NewClient(ctx, option.WithCredentialsFile("google-cloud-storage-credentials.json"))

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})

	if err != nil {
		panic("storage client error: " + err.Error())
	}

	return client
}
