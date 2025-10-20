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

func (repo *Repository) CreateProvince(ctx context.Context, province model.Province) (model.Province, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Create(&province).Error
	if err != nil {
		logger.Error(ctx, "Error CreateProvince", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "province", "repo"},
		})

		return province, err
	}

	return province, nil
}

func (repo *Repository) GetProvinces(ctx context.Context, query request.GetProvinceQuery) ([]model.Province, error) {
	var res []model.Province
	statement := repo.db.WithContext(ctx).Model(&res)

	if query.Search != "" {
		querySearch := fmt.Sprintf("%s%s%s", "%", query.Search, "%")
		statement = statement.Where("name ILIKE ?", querySearch)
	}

	order := query.GetOrderQuery()
	if order != "" {
		statement = statement.Order(order)
	}

	err := statement.Limit(int(query.Limit)).
		Offset(int(query.GetOffset())).
		Find(&res).Error

	if err != nil {
		logger.Error(ctx, "Error GetProvince", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "province", "repo"},
		})

		return res, err
	}

	return res, nil
}

func (repo *Repository) GetProvinceTotal(ctx context.Context, query request.GetProvinceQuery) (uint, error) {

	var total int64
	var province model.Province
	statement := repo.db.WithContext(ctx).Model(&province)

	if query.Search != "" {
		querySearch := fmt.Sprintf("%s%s%s", "%", query.Search, "%")
		statement = statement.Where("name ILIKE ?", querySearch)
	}

	order := query.GetOrderQuery()
	if order != "" {
		statement = statement.Order(order)
	}

	err := statement.Count(&total).Error

	if err != nil {
		logger.Error(ctx, "Error GetProvinceTotal", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "province", "repo"},
		})

		return uint(total), err
	}

	return uint(total), nil
}

func (repo *Repository) UpdateProvince(ctx context.Context, province model.Province) (model.Province, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Save(&province).Error
	if err != nil {
		logger.Error(ctx, "Error UpdateProvince", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "province", "repo"},
		})

		return province, err
	}

	return province, nil
}

func (repo *Repository) UpdateProvinceAddTotalCity(ctx context.Context, provinceID uint, add int) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Exec("UPDATE province SET total_city = (SELECT total_city FROM province WHERE id = ?)+? WHERE id = ?", provinceID, add, provinceID).Error
	if err != nil {
		logger.Error(ctx, "Error UpdateProvinceAddTotalCity", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "province", "repo"},
		})

		return err
	}

	return nil
}

func (repo *Repository) UpdateProvinceAddTotalDistrict(ctx context.Context, provinceID uint, add int) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Exec("UPDATE province SET total_district = (SELECT total_district FROM province WHERE id = ?)+? WHERE id = ?", provinceID, add, provinceID).Error
	if err != nil {
		logger.Error(ctx, "Error UpdateProvinceAddTotalDistrict", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "province", "repo"},
		})

		return err
	}

	return nil
}

func (repo *Repository) UpdateProvinceAddTotalVillage(ctx context.Context, provinceID uint, add int) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Exec("UPDATE province SET total_village = (SELECT total_village FROM province WHERE id = ?)+? WHERE id = ?", provinceID, add, provinceID).Error
	if err != nil {
		logger.Error(ctx, "Error UpdateProvinceAddTotalVillage", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "province", "repo"},
		})

		return err
	}

	return nil
}

func (repo *Repository) GetProvince(ctx context.Context, query request.GetProvinceQuery) (province model.Province, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	statement := tx.WithContext(ctx).Model(&province)

	if query.ID != 0 {
		statement = statement.Where("id = ?", query.ID)
	}

	if query.Code != "" {
		statement = statement.Where("code = ?", query.Code)
	}

	if query.IsActiveFlag != "" {
		statement = statement.Where("is_active_flag = ?", query.IsActiveFlag)
	}

	if len(query.Preloads) > 0 {
		for _, p := range query.Preloads {
			statement = statement.Preload(p)
		}
	}

	err = statement.First(&province).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return province, lib.ErrorNotFound
	}

	if err != nil {
		logger.Error(ctx, "Error GetProvinceByID", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "province", "repo"},
		})

		return province, err
	}

	return province, nil
}

func (repo *Repository) DeleteProvince(ctx context.Context, ID uint) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var province model.Province
	err := tx.WithContext(ctx).Delete(&province, ID).Error
	if err != nil {
		logger.Error(ctx, "Error DeleteProvince", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "province", "repo"},
		})

		return err
	}

	return nil
}
