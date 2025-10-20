package repository

import (
	"context"
	"errors"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/logger"
	"gorm.io/gorm"
)

type Repository struct {
	db          *lib.Database
	smtpClient  *lib.SMTPClient
	redisClient *lib.RedisConnection
}

func NewRepository(db *lib.Database, smtpClient *lib.SMTPClient, redisClient *lib.RedisConnection) Repository {
	return Repository{db: db, smtpClient: smtpClient, redisClient: redisClient}
}

func (repo *Repository) Transaction(ctx context.Context, fn func(context.Context) error) error {
	trx := repo.db.Begin()

	ctx = context.WithValue(ctx, "Trx", &lib.Database{DB: trx})
	if err := fn(ctx); err != nil {
		trx.Rollback()
		return err
	}

	return trx.Commit().Error
}

func (repo *Repository) SendEmail(ctx context.Context, req lib.SMTPRequest) error {
	return repo.smtpClient.Send(ctx, req)
}

func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func LogError(ctx context.Context, message string, err error) {
	logger.Error(ctx, message, map[string]interface{}{
		"error": err,
		"tags":  []string{"gorm"},
	})
}

func LogWarn(ctx context.Context, message string) {
	logger.Warn(ctx, message, map[string]interface{}{
		"tags": []string{"gorm"},
	})
}
