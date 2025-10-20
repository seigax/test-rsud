package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/logger"
)

type Local struct {
	Directory string // relative public
}

func (m *Local) UploadFile(ctx context.Context, bucketName, fileName, contentType string, file io.Reader) error {
	path := filepath.Join(m.Directory, bucketName, fileName)
	os.MkdirAll(filepath.Dir(path), os.ModePerm)

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		logger.Error(ctx, "failed open local file", map[string]interface{}{
			"error": err,
			"tags":  []string{"storage", "local"},
		})

		return err
	}

	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		logger.Error(ctx, "failed copy local file", map[string]interface{}{
			"error": err,
			"tags":  []string{"storage", "local"},
		})

		return err
	}

	if fileInfo, err := os.Stat(path); fileInfo.Size() > 10000000 {
		if err != nil {
			return errors.New("file not found")
		}

		return errors.New("file too large")
	}

	return nil
}

func (m *Local) GetFileTemporaryURL(ctx context.Context, bucketName, filename string) (string, error) {
	return fmt.Sprintf("%s/%s", viper.GetString("CDN_BASE_URL"), filename), nil
}

func (m *Local) GetObject(_ context.Context, _, filename string) (io.Reader, error) {
	return os.Open(fmt.Sprintf("%s", filename))
}

func (m *Local) FGetObject(ctx context.Context, bucketName, filename, destination string) error {
	return nil
}

func (m *Local) FPutObject(ctx context.Context, bucketName, filename, source string) error {
	return nil
}

func (m *Local) RemoveFile(ctx context.Context, bucketName, filename string) error {
	return nil
}
