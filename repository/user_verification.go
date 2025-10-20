package repository

import (
	"context"
	"errors"
	"time"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/logger"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gorm.io/gorm"
)

func (repo *Repository) CreateUserVerification(ctx context.Context, user model.UserVerification) (model.UserVerification, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Create(&user).Error
	if err != nil {
		logger.Error(ctx, "Error CreateUserVerification", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "userPhone", "repo"},
		})

		return user, err
	}

	return user, nil
}

func (repo *Repository) GetUserVerificationByUserId(ctx context.Context, UserId uint) (model.UserVerification, error) {
	var userVerifaction model.UserVerification
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Table("user_verification").Select("*").
		Where("user_id = ?", UserId).
		Where("verification_status_flag = ?", false).
		Where("expired_at >= ?", time.Now()).
		Order("created_at DESC").
		First(&userVerifaction).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return userVerifaction, lib.OtpExpired
	}

	if err != nil {
		logger.Error(ctx, "Error GetUserVerificationByUserId", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "city", "repo"},
		})

		return userVerifaction, err
	}

	return userVerifaction, nil
}

func (repo *Repository) UpdateVerificationStatusToTrue(ctx context.Context, ID uint) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Table("user_verification").Where("id = ?", ID).Update("verification_status_flag", true).Error

	if err != nil {
		logger.Error(ctx, "Error UpdateVerificationStatusToTrue", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "userPhone", "repo"},
		})

		return err
	}

	return nil
}
