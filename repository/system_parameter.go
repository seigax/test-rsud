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

func (repo *Repository) CreateSystemParameter(ctx context.Context, SystemParameter model.SystemParameter) (model.SystemParameter, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Create(&SystemParameter).Error
	if err != nil {
		logger.Error(ctx, "Error CreateSystemParameter", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "SystemParameter", "repo"},
		})

		return SystemParameter, err
	}

	return SystemParameter, nil
}

func (repo *Repository) GetSystemParameters(ctx context.Context, query request.GetSystemParameterQuery) ([]model.SystemParameter, error) {

	var res []model.SystemParameter
	statement := repo.db.WithContext(ctx).Model(&res)

	if query.Search != "" {
		querySearch := fmt.Sprintf("%s%s%s", "%", query.Search, "%")
		statement = statement.Where("parameter_name ILIKE ?", querySearch)
	}

	order := query.GetOrderQuery()
	if order != "" {
		statement = statement.Order(order)
	}

	err := statement.Limit(int(query.Limit)).
		Offset(int(query.GetOffset())).
		Find(&res).Error

	if err != nil {
		logger.Error(ctx, "Error GetSystemParameter", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "SystemParameter", "repo"},
		})

		return res, err
	}

	return res, nil
}

func (repo *Repository) GetSystemParameterTotal(ctx context.Context, query request.GetSystemParameterQuery) (uint, error) {

	var total int64
	var SystemParameter model.SystemParameter
	statement := repo.db.WithContext(ctx).Model(&SystemParameter)

	if query.Search != "" {
		querySearch := fmt.Sprintf("%s%s%s", "%", query.Search, "%")
		statement = statement.Where("parameter_name ILIKE ?", querySearch)
	}

	order := query.GetOrderQuery()
	if order != "" {
		statement = statement.Order(order)
	}

	err := statement.Count(&total).Error

	if err != nil {
		logger.Error(ctx, "Error GetSystemParameterTotal", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "SystemParameter", "repo"},
		})

		return uint(total), err
	}

	return uint(total), nil
}

func (repo *Repository) UpdateSystemParameter(ctx context.Context, SystemParameter model.SystemParameter) (model.SystemParameter, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Save(&SystemParameter).Error
	if err != nil {
		logger.Error(ctx, "Error UpdateSystemParameter", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "SystemParameter", "repo"},
		})

		return SystemParameter, err
	}

	return SystemParameter, nil
}

func (repo *Repository) GetSystemParameter(ctx context.Context, query request.GetSystemParameterQuery) (systemParameter model.SystemParameter, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	statement := tx.WithContext(ctx).Model(&systemParameter)

	if query.ID != 0 {
		statement = statement.Where("id = ?", query.ID)
	}

	if query.Code != "" {
		statement = statement.Where("code = ?", query.Code)
	}

	if query.ParameterName != "" {
		statement = statement.Where("parameter_name = ?", query.ParameterName)
	}

	if query.DataType != "" {
		statement = statement.Where("data_type = ?", query.DataType)
	}

	if query.Message != "" {
		statement = statement.Where("message = ?", query.Message)
	}

	if query.IsActiveFlag != "" {
		statement = statement.Where("is_active_flag = ?", query.IsActiveFlag)
	}

	if len(query.Preloads) > 0 {
		for _, p := range query.Preloads {
			statement = statement.Preload(p)
		}
	}

	err = statement.First(&systemParameter).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return systemParameter, lib.ErrorNotFound
	}

	if err != nil {
		logger.Error(ctx, "Error GetSystemParameterByID", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "SystemParameter", "repo"},
		})

		return systemParameter, err
	}

	return systemParameter, nil
}

func (repo *Repository) DeleteSystemParameter(ctx context.Context, ID uint) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var SystemParameter model.SystemParameter
	err := tx.WithContext(ctx).Delete(&SystemParameter, ID).Error
	if err != nil {
		logger.Error(ctx, "Error DeleteSystemParameter", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "SystemParameter", "repo"},
		})

		return err
	}

	return nil
}
