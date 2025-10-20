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

func (repo *Repository) CreateDistrict(ctx context.Context, district model.District) (model.District, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Create(&district).Error
	if err != nil {
		logger.Error(ctx, "Error CreateDistrict", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "district", "repo"},
		})

		return district, err
	}

	return district, nil
}

func (repo *Repository) GetDistricts(ctx context.Context, query request.GetDistrictQuery) ([]model.District, error) {

	var res []model.District
	statement := repo.db.WithContext(ctx).Model(&res)

	if query.Search != "" {
		querySearch := fmt.Sprintf("%s%s%s", "%", query.Search, "%")
		statement = statement.Where("name ILIKE ?", querySearch)
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

	if query.CityID != 0 {
		statement = statement.Where("city_id = ?", query.CityID)
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
		logger.Error(ctx, "Error GetDistrict", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "district", "repo"},
		})

		return res, err
	}

	return res, nil
}

func (repo *Repository) GetDistrictTotal(ctx context.Context, query request.GetDistrictQuery) (uint, error) {

	var total int64
	var District model.District
	statement := repo.db.WithContext(ctx).Model(&District)

	if query.Search != "" {
		querySearch := fmt.Sprintf("%s%s%s", "%", query.Search, "%")
		statement = statement.Where("name ILIKE ?", querySearch)
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

	if query.CityID != 0 {
		statement = statement.Where("city_id = ?", query.CityID)
	}

	if query.IsActiveFlag != "" {
		statement = statement.Where("is_active_flag = ?", query.IsActiveFlag)
	}

	err := statement.Count(&total).Error

	if err != nil {
		logger.Error(ctx, "Error GetDistrictTotal", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "district", "repo"},
		})

		return uint(total), err
	}

	return uint(total), nil
}

func (repo *Repository) UpdateDistrict(ctx context.Context, district model.District) (model.District, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Save(&district).Error
	if err != nil {
		logger.Error(ctx, "Error UpdateDistrict", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "district", "repo"},
		})

		return district, err
	}

	return district, nil
}

func (repo *Repository) GetDistrict(ctx context.Context, query request.GetDistrictQuery) (district model.District, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	statement := tx.WithContext(ctx).Model(&district)

	if query.ID != 0 {
		statement = statement.Where("id = ?", query.ID)
	}

	if query.Code != "" {
		statement = statement.Where("code = ?", query.Code)
	}

	if query.CityID != 0 {
		statement = statement.Where("city_id = ?", query.CityID)
	}

	if query.IsActiveFlag != "" {
		statement = statement.Where("is_active_flag = ?", query.IsActiveFlag)
	}

	if len(query.Preloads) > 0 {
		for _, p := range query.Preloads {
			statement = statement.Preload(p)
		}
	}

	err = statement.First(&district).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return district, lib.ErrorNotFound
	}

	if err != nil {
		logger.Error(ctx, "Error GetDistrict", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "district", "repo"},
		})

		return district, err
	}

	return district, nil
}

func (repo *Repository) DeleteDistrict(ctx context.Context, ID uint) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var district model.District
	err := tx.WithContext(ctx).Delete(&district, ID).Error
	if err != nil {
		logger.Error(ctx, "Error DeleteDistrict", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "district", "repo"},
		})

		return err
	}

	return nil
}
