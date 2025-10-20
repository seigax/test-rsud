package repository

import (
	"context"
	"time"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/logger"
)

func (repo *Repository) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	err := repo.redisClient.Client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		logger.Error(ctx, "Redis Set Error", map[string]interface{}{
			"error": err,
			"key":   key,
			"value": value,
			"tags":  []string{"redis", "repo"},
		})

		return err
	}
	return nil
}

func (repo *Repository) Get(ctx context.Context, key string) (string, error) {
	value, err := repo.redisClient.Client.Get(ctx, key).Result()
	if err != nil {
		logger.Error(ctx, "Redis Get Error", map[string]interface{}{
			"error": err,
			"key":   key,
			"tags":  []string{"redis", "repo"},
		})

		return "", err
	}
	return value, nil
}

func (repo *Repository) Delete(ctx context.Context, key string) error {
	err := repo.redisClient.Client.Del(ctx, key).Err()
	if err != nil {
		logger.Error(ctx, "Redis Delete Error", map[string]interface{}{
			"error": err,
			"key":   key,
			"tags":  []string{"redis", "repo"},
		})

		return err
	}
	return nil
}
