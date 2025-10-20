package config

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
)

func NewRedisConnection() *lib.RedisConnection {
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", viper.GetString("REDIS_HOST"), viper.GetString("REDIS_PORT")),
		Password:     viper.GetString("REDIS_PASSWORD"),
		Username:     viper.GetString("REDIS_USERNAME"),
		DialTimeout:  viper.GetDuration("REDIS_DIAL_TIMEOUT"),
		ReadTimeout:  viper.GetDuration("REDIS_READ_TIMEOUT"),
		WriteTimeout: viper.GetDuration("REDIS_WRITE_TIMEOUT"),
		DB:           0,
	})
	return &lib.RedisConnection{
		Client: client,
	}
}
