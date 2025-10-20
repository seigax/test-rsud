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

func (repo *Repository) CreateVillage(ctx context.Context, village model.Village) (model.Village, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Create(&village).Error
	if err != nil {
		logger.Error(ctx, "Error CreateVillage", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "village", "repo"},
		})

		return village, err
	}

	return village, nil
}

func (repo *Repository) GetVillages(ctx context.Context, query request.GetVillageQuery) ([]model.Village, error) {

	var res []model.Village
	statement := repo.db.WithContext(ctx).Model(&res).
		Where("is_active_flag = ?", "Y")

	if query.Search != "" {
		querySearch := fmt.Sprintf("%s%s%s", "%", query.Search, "%")
		statement = statement.Where("name ILIKE ?", querySearch)
	}

	order := query.GetOrderQuery()
	if order != "" {
		statement = statement.Order(order)
	}

	statement = statement.Where("district_id = ?", query.DistrictID)

	err := statement.Limit(int(query.Limit)).
		Offset(int(query.GetOffset())).
		Find(&res).Error

	if err != nil {
		logger.Error(ctx, "Error GetVillage", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "village", "repo"},
		})

		return res, err
	}

	return res, nil
}

func (repo *Repository) GetVillageTotal(ctx context.Context, query request.GetVillageQuery) (uint, error) {

	var total int64
	var village model.Village
	statement := repo.db.WithContext(ctx).Model(&village).
		Where("is_active_flag = ?", "Y")

	if query.Search != "" {
		querySearch := fmt.Sprintf("%s%s%s", "%", query.Search, "%")
		statement = statement.Where("name ILIKE ?", querySearch)
	}

	order := query.GetOrderQuery()
	if order != "" {
		statement = statement.Order(order)
	}

	statement = statement.Where("district_id = ?", query.DistrictID)

	err := statement.Count(&total).Error

	if err != nil {
		logger.Error(ctx, "Error GetVillageTotal", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "village", "repo"},
		})

		return uint(total), err
	}

	return uint(total), nil
}

func (repo *Repository) UpdateVillage(ctx context.Context, village model.Village) (model.Village, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Save(&village).Error
	if err != nil {
		logger.Error(ctx, "Error UpdateVillage", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "village", "repo"},
		})

		return village, err
	}

	return village, nil
}

func (repo *Repository) GetVillageByID(ctx context.Context, ID uint) (village model.Village, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err = tx.WithContext(ctx).First(&village, ID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return village, lib.ErrorNotFound
	}

	if err != nil {
		logger.Error(ctx, "Error GetVillageByID", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "village", "repo"},
		})

		return village, err
	}

	return village, nil
}

func (repo *Repository) GetVillagesByDistrictID(ctx context.Context, districtID uint) (villages []model.Village, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err = tx.WithContext(ctx).
		Where("district_id = ?", districtID).
		Find(&villages).Error

	if err != nil {
		logger.Error(ctx, "Error GetVillagesByDistrictID", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "village", "repo"},
		})

		return
	}

	return
}

func (repo *Repository) DeleteVillage(ctx context.Context, ID uint) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var village model.Village
	err := tx.WithContext(ctx).Delete(&village, ID).Error
	if err != nil {
		logger.Error(ctx, "Error DeleteVillage", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "village", "repo"},
		})

		return err
	}

	return nil
}
