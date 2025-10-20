package repository

import (
	"context"
	"errors"
	"fmt"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/logger"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
	"gorm.io/gorm"
)

func (repo *Repository) CreateErrorMessage(ctx context.Context, errorMessage model.ErrorMessage) (model.ErrorMessage, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Create(&errorMessage).Error
	if err != nil {
		logger.Error(ctx, "Error CreateErrorMessage", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "errorMessage", "repo"},
		})

		return errorMessage, err
	}

	return errorMessage, nil
}

func (repo *Repository) GetErrorMessages(ctx context.Context, query request.GetErrorMessageQuery) ([]model.ErrorMessage, error) {

	var res []model.ErrorMessage
	statement := repo.db.WithContext(ctx).Model(&res)

	if query.Search != "" {
		querySearch := fmt.Sprintf("%s%s%s", "%", query.Search, "%")
		statement = statement.Where("message ILIKE ?", querySearch)
	}

	order := query.GetOrderQuery()
	if order != "" {
		statement = statement.Order(order)
	}

	if query.ID != 0 {
		statement = statement.Where("id = ?", query.ID)
	}

	if query.Code != "" {
		statement = statement.Where("code = ?", query.Code)
	}

	if query.Type != "" {
		statement = statement.Where("type = ?", query.Type)
	}

	if query.ApplicationName != "" {
		statement = statement.Where("application_name = ?", query.ApplicationName)
	}

	if query.IsActiveFlag != "" {
		statement = statement.Where("is_active_flag = ?", query.IsActiveFlag)
	}

	if len(query.Preloads) > 0 {
		for _, p := range query.Preloads {
			statement = statement.Preload(p)
		}
	}

	err := statement.Limit(int(query.Limit)).
		Offset(int(query.GetOffset())).
		Find(&res).Error

	if err != nil {
		logger.Error(ctx, "Error GetErrorMessage", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "errorMessage", "repo"},
		})

		return res, err
	}

	return res, nil
}

func (repo *Repository) GetErrorMessageTotal(ctx context.Context, query request.GetErrorMessageQuery) (uint, error) {

	var total int64
	var errorMessage model.ErrorMessage
	statement := repo.db.WithContext(ctx).Model(&errorMessage)

	if query.Search != "" {
		querySearch := fmt.Sprintf("%s%s%s", "%", query.Search, "%")
		statement = statement.Where("message ILIKE ?", querySearch)
	}

	order := query.GetOrderQuery()
	if order != "" {
		statement = statement.Order(order)
	}

	if query.ID != 0 {
		statement = statement.Where("id = ?", query.ID)
	}

	if query.Code != "" {
		statement = statement.Where("code = ?", query.Code)
	}

	if query.Type != "" {
		statement = statement.Where("type = ?", query.Type)
	}

	if query.ApplicationName != "" {
		statement = statement.Where("application_name = ?", query.ApplicationName)
	}

	if query.IsActiveFlag != "" {
		statement = statement.Where("is_active_flag = ?", query.IsActiveFlag)
	}

	err := statement.Count(&total).Error

	if err != nil {
		logger.Error(ctx, "Error GetErrorMessageTotal", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "errorMessage", "repo"},
		})

		return uint(total), err
	}

	return uint(total), nil
}

func (repo *Repository) UpdateErrorMessage(ctx context.Context, errorMessage model.ErrorMessage) (model.ErrorMessage, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Save(&errorMessage).Error
	if err != nil {
		logger.Error(ctx, "Error UpdateErrorMessage", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "errorMessage", "repo"},
		})

		return errorMessage, err
	}

	return errorMessage, nil
}

func (repo *Repository) GetErrorMessage(ctx context.Context, query request.GetErrorMessageQuery) (errorMessage model.ErrorMessage, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	statement := tx.WithContext(ctx).Model(&errorMessage)

	if query.ID != 0 {
		statement = statement.Where("id = ?", query.ID)
	}

	if query.Code != "" {
		statement = statement.Where("code = ?", query.Code)
	}

	if query.Type != "" {
		statement = statement.Where("type = ?", query.Type)
	}

	if query.ApplicationName != "" {
		statement = statement.Where("application_name = ?", query.ApplicationName)
	}

	if query.IsActiveFlag != "" {
		statement = statement.Where("is_active_flag = ?", query.IsActiveFlag)
	}

	if len(query.Preloads) > 0 {
		for _, p := range query.Preloads {
			statement = statement.Preload(p)
		}
	}

	err = statement.First(&errorMessage).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorMessage, lib.ErrorNotFound
	}

	if err != nil {
		logger.Error(ctx, "Error GetErrorMessage", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "errorMessage", "repo"},
		})

		return errorMessage, err
	}

	return errorMessage, nil
}

func (repo *Repository) DeleteErrorMessage(ctx context.Context, ID uint) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var errorMessage model.ErrorMessage
	err := tx.WithContext(ctx).Delete(&errorMessage, ID).Error
	if err != nil {
		logger.Error(ctx, "Error DeleteErrorMessage", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "errorMessage", "repo"},
		})

		return err
	}

	return nil
}
