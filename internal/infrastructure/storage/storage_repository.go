package storage

import (
	"context"
	"io"
	"mime/multipart"
	storageDomain "my-story-time-api/internal/domain/storage"

	"cloud.google.com/go/storage"
)

type storageRepository struct {
	client *storage.Client
}

func NewStorageRepository(client *storage.Client) storageDomain.StorageRepository {
	return &storageRepository{client}
}

func (r *storageRepository) Upload(file *multipart.FileHeader, objectName string, bucketName string) (*storage.Writer, error) {
	ctx := context.Background()

	src, err := file.Open()

	if err != nil {
		return nil, err
	}

	defer src.Close()

	bucket := r.client.Bucket(bucketName)
	object := bucket.Object(objectName)

	wc := object.NewWriter(ctx)

	if _, err := io.Copy(wc, src); err != nil {
		return nil, err
	}
	if err := wc.Close(); err != nil {
		return nil, err
	}

	return wc, nil
}
