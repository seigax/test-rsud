package repository

import (
	"context"
	"errors"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/logger"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gorm.io/gorm"
)

func (repo *Repository) CreateSession(ctx context.Context, session model.Session) (model.Session, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Create(&session).Error
	if err != nil {
		logger.Error(ctx, "Error CreateSession", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "session", "repo"},
		})

		return session, err
	}

	return session, nil
}

func (repo *Repository) UpdateSession(ctx context.Context, session model.Session) (model.Session, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Save(&session).Error
	if err != nil {
		logger.Error(ctx, "Error UpdateSession", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "session", "repo"},
		})

		return session, err
	}

	return session, nil
}

func (repo *Repository) GetSession(ctx context.Context, session model.Session) (model.Session, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	statement := tx.WithContext(ctx).Model(&model.Session{})

	if session.ID != 0 {
		statement = statement.Where("id = ?", session.ID)
	}

	if session.Token != "" {
		statement = statement.Where("token = ?", session.Token)
	}

	err := statement.First(&session).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return session, lib.ErrorNotFound
	}

	if err != nil {
		logger.Error(ctx, "Error GetSessionByID", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "session", "repo"},
		})

		return session, err
	}

	return session, nil
}

func (repo *Repository) DeleteSession(ctx context.Context, ID uint) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var session model.Session
	err := tx.WithContext(ctx).Delete(&session, ID).Error
	if err != nil {
		logger.Error(ctx, "Error DeleteSession", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "session", "repo"},
		})

		return err
	}

	return nil
}
