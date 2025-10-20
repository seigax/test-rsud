package storage

import (
	"context"
	"io"
)

type Storage interface {
	UploadFile(ctx context.Context, bucketName, fileName, contentType string, file io.Reader) error
	GetFileTemporaryURL(ctx context.Context, bucketName, filename string) (string, error)
	GetObject(ctx context.Context, bucketName, filename string) (io.Reader, error)
	FGetObject(ctx context.Context, bucketName, filename, destination string) error
	FPutObject(ctx context.Context, bucketName, filename, source string) error
	RemoveFile(ctx context.Context, bucketName, pathFilename string) error
}
