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

func (repo *Repository) CreateCity(ctx context.Context, City model.City) (model.City, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Create(&City).Error
	if err != nil {
		logger.Error(ctx, "Error CreateCity", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "city", "repo"},
		})

		return City, err
	}

	return City, nil
}

func (repo *Repository) GetCitys(ctx context.Context, query request.GetCityQuery) ([]model.City, error) {

	var res []model.City
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

	if query.ProvinceID != 0 {
		statement = statement.Where("province_id = ?", query.ProvinceID)
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
		logger.Error(ctx, "Error GetCity", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "city", "repo"},
		})

		return res, err
	}

	return res, nil
}

func (repo *Repository) GetCityTotal(ctx context.Context, query request.GetCityQuery) (uint, error) {

	var total int64
	var City model.City
	statement := repo.db.WithContext(ctx).Model(&City)

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

	if query.ProvinceID != 0 {
		statement = statement.Where("province_id = ?", query.ProvinceID)
	}

	if query.IsActiveFlag != "" {
		statement = statement.Where("is_active_flag = ?", query.IsActiveFlag)
	}

	err := statement.Count(&total).Error

	if err != nil {
		logger.Error(ctx, "Error GetCityTotal", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "city", "repo"},
		})

		return uint(total), err
	}

	return uint(total), nil
}

func (repo *Repository) UpdateCity(ctx context.Context, city model.City) (model.City, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Save(&city).Error
	if err != nil {
		logger.Error(ctx, "Error UpdateCity", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "City", "repo"},
		})

		return city, err
	}

	return city, nil
}

func (repo *Repository) GetCity(ctx context.Context, query request.GetCityQuery) (city model.City, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	statement := tx.WithContext(ctx).Model(&city)

	if query.ID != 0 {
		statement = statement.Where("id = ?", query.ID)
	}

	if query.Code != "" {
		statement = statement.Where("code = ?", query.Code)
	}

	if query.ProvinceID != 0 {
		statement = statement.Where("province_id = ?", query.ProvinceID)
	}

	if query.IsActiveFlag != "" {
		statement = statement.Where("is_active_flag = ?", query.IsActiveFlag)
	}

	if len(query.Preloads) > 0 {
		for _, p := range query.Preloads {
			statement = statement.Preload(p)
		}
	}

	err = statement.First(&city).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return city, lib.ErrorNotFound
	}

	if err != nil {
		logger.Error(ctx, "Error GetCity", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "city", "repo"},
		})

		return city, err
	}

	return city, nil
}

func (repo *Repository) DeleteCity(ctx context.Context, ID uint) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var City model.City
	err := tx.WithContext(ctx).Delete(&City, ID).Error
	if err != nil {
		logger.Error(ctx, "Error DeleteCity", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "city", "repo"},
		})

		return err
	}

	return nil
}
