package lib

import (
	"github.com/go-redis/redis/v8"
)

type RedisConnection struct {
	Client *redis.Client
}
