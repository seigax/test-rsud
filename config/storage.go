package config

import (
	"github.com/spf13/viper"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/storage"
)

func NewLocalStorage() *storage.Local {
	return &storage.Local{Directory: "public"}
}

func NewStorage() storage.Storage {
	if !viper.GetBool("FORCE_LOCAL_STORAGE") {
		minioClient := NewMinioClient()
		return &storage.Minio{Client: minioClient}
	}

	return NewLocalStorage()
}
