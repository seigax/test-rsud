package storage

import (
	"context"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/viper"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/logger"
)

type Minio struct {
	Client *lib.MinioClient
}

func (m *Minio) UploadFile(ctx context.Context, bucketName, fileName, contentType string, file io.Reader) error {
	_, err := m.Client.PutObject(ctx, bucketName, fileName, file, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		logger.Error(ctx, "failed upload file to minio", map[string]interface{}{
			"error": err,
			"tags":  []string{"storage", "minio"},
		})
	}

	return err
}

func (m *Minio) GetFileTemporaryURL(ctx context.Context, bucketName, filename string) (string, error) {
	// Set request parameters for content-disposition.
	reqParams := make(url.Values)

	// Generates a presigned url which expires in a day.
	presignedURL, err := m.Client.PresignedGetObject(ctx, bucketName, filename, time.Second*24*60*60, reqParams)
	if err != nil {
		logger.Error(ctx, "failed get file temporary url", map[string]interface{}{
			"error": err,
			"tags":  []string{"storage", "minio"},
		})

		return "", err
	}

	baseURL := viper.GetString("CDN_BASE_DNS")
	if baseURL == "" {
		baseURL = viper.GetString("CDN_BASE_URL")
	}

	var result string
	if baseURL == "https://minio.asaba.co.id" {
		result = baseURL + presignedURL.Path
	} else {
		result = baseURL + presignedURL.Path + "?" + presignedURL.RawQuery
	}

	return result, nil
}

func (m *Minio) GetObject(ctx context.Context, bucketName, filename string) (io.Reader, error) {
	return m.Client.GetObject(ctx, bucketName, filename, minio.GetObjectOptions{})
}

func (m *Minio) FGetObject(ctx context.Context, bucketName, filename, destination string) error {
	return m.Client.FGetObject(ctx, bucketName, filename, destination, minio.GetObjectOptions{})
}

func (m *Minio) FPutObject(ctx context.Context, bucketName, filename, source string) error {
	_, err := m.Client.FPutObject(ctx, bucketName, filename, source, minio.PutObjectOptions{})
	return err
}

func (m *Minio) RemoveFile(ctx context.Context, bucketName, pathFilename string) error {
	opts := minio.RemoveObjectOptions{GovernanceBypass: true}
	err := m.Client.RemoveObject(ctx, bucketName, pathFilename, opts)
	if err != nil {
		logger.Error(ctx, "failed upload file to minio", map[string]interface{}{
			"error": err,
			"tags":  []string{"storage", "minio"},
		})
	}

	return err
}
