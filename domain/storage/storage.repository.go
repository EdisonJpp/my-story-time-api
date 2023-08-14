package storage

import (
	"cloud.google.com/go/storage"
	"mime/multipart"
)

type StorageRepository interface {
	Upload(file *multipart.FileHeader, fileName string, bucketName string) (*storage.Writer, error)
}
