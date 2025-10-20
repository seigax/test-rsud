package config

import (
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
)

func NewMinioClient() *lib.MinioClient {
	endpoint := viper.GetString("MINIO_ENDPOINT")
	username := viper.GetString("MINIO_USERNAME")
	password := viper.GetString("MINIO_PASSWORD")
	useSSL, err := strconv.ParseBool(viper.GetString("MINIO_SSL"))
	if err != nil {
		panic(err)
	}

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(username, password, ""),
		Secure: useSSL,
	})

	if err != nil {
		panic(err)
	}

	return &lib.MinioClient{Client: client}
}
