package repository

import (
	"context"
	"errors"
	"time"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/logger"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gorm.io/gorm"
)

func (repo *Repository) CreateUserPhone(ctx context.Context, userPhone model.UserPhone) (model.UserPhone, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Create(&userPhone).Error
	if err != nil {
		logger.Error(ctx, "Error CreateUserPhone", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "userPhone", "repo"},
		})

		return userPhone, err
	}

	return userPhone, nil
}

func (repo *Repository) GetUserPhones(ctx context.Context, userID uint) ([]model.UserPhone, error) {
	var res []model.UserPhone
	statement := repo.db.WithContext(ctx).Model(&res)

	err := statement.
		Find(&res).Error

	if err != nil {
		logger.Error(ctx, "Error GetUserPhone", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "userPhone", "repo"},
		})

		return res, err
	}

	return res, nil
}

func (repo *Repository) UpdateUserPhone(ctx context.Context, userPhone model.UserPhone) (model.UserPhone, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Save(&userPhone).Error
	if err != nil {
		logger.Error(ctx, "Error UpdateUserPhone", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "userPhone", "repo"},
		})

		return userPhone, err
	}

	return userPhone, nil
}

func (repo *Repository) GetUserPhoneByID(ctx context.Context, ID uint) (userPhone model.UserPhone, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err = tx.WithContext(ctx).First(&userPhone, ID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return userPhone, lib.ErrorNotFound
	}

	if err != nil {
		logger.Error(ctx, "Error GetUserPhoneByID", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "userPhone", "repo"},
		})

		return userPhone, err
	}

	return userPhone, nil
}

func (repo *Repository) GetUserPhone(ctx context.Context, phoneNumber string) (userPhone model.UserPhone, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err = tx.WithContext(ctx).Where("phone_number = ?", phoneNumber).First(&userPhone).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return userPhone, lib.ErrorNotFound
	}

	if err != nil {
		logger.Error(ctx, "Error GetUserPhoneByID", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "userPhone", "repo"},
		})

		return userPhone, err
	}

	return userPhone, nil
}

func (repo *Repository) CheckPhoneNumberIsUsed(ctx context.Context, phoneNumber string) (bool, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var userPhone model.UserPhone

	err := tx.WithContext(ctx).Where("phone_number = ?", phoneNumber).First(&userPhone).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, lib.ErrorNotFound
	}

	if err != nil {
		logger.Error(ctx, "Error CheckPhoneNumberIsUsed", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "userPhone", "repo"},
		})

		return false, err
	}

	return true, nil
}

func (repo *Repository) DeleteUserPhone(ctx context.Context, ID uint) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var userPhone model.UserPhone
	err := tx.WithContext(ctx).Delete(&userPhone, ID).Error
	if err != nil {
		logger.Error(ctx, "Error DeleteUserPhone", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "userPhone", "repo"},
		})

		return err
	}

	return nil
}

func (repo *Repository) GetAllUserPhones(ctx context.Context, userID uint) ([]model.UserPhone, error) {
	var res []model.UserPhone
	statement := repo.db.WithContext(ctx).Model(&res).Where("user_id = ?", userID)

	err := statement.
		Find(&res).Error

	if err != nil {
		logger.Error(ctx, "Error GetUserPhone", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "userPhone", "repo"},
		})

		return res, err
	}

	return res, nil
}

func (repo *Repository) ChangeActiveUserPhone(ctx context.Context, req request.ReqUpdatePhone) error {
	var res []model.UserPhone
	err := repo.db.WithContext(ctx).Model(&res).
		Where("user_id = ?", req.UserId).
		Where("id = ?", req.Id).
		Updates(map[string]interface{}{"updated_by": req.UserId, "updated_at": time.Now(), "is_active_flag": "Y"}).Error

	if err != nil {
		logger.Error(ctx, "Error ChangeActiveUserPhone", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "userPhone", "repo"},
		})

		return err
	}

	err = repo.db.WithContext(ctx).Model(&res).
		Where("user_id = ?", req.UserId).
		Where("id != ?", req.Id).
		Updates(map[string]interface{}{"updated_by": req.UserId, "updated_at": time.Now(), "is_active_flag": "N"}).Error

	if err != nil {
		logger.Error(ctx, "Error ChangeActiveUserPhone", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "userPhone", "repo"},
		})

		return err
	}

	return nil
}
