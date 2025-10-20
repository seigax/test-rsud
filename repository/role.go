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

func (repo *Repository) CreateRole(ctx context.Context, role model.Role) (model.Role, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Create(&role).Error
	if err != nil {
		logger.Error(ctx, "Error CreateRole", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "role", "repo"},
		})

		return role, err
	}

	return role, nil
}

func (repo *Repository) GetRoles(ctx context.Context, query request.GetRoleQuery) ([]model.Role, error) {
	var res []model.Role
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

	if query.Type != "" {
		statement = statement.Where("type = ?", query.Type)
	}

	if query.Platform != "" {
		statement = statement.Where("platform = ?", query.Platform)
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
		logger.Error(ctx, "Error GetRoles", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "role", "repo"},
		})

		return res, err
	}

	return res, nil
}

func (repo *Repository) GetRoleTotal(ctx context.Context, query request.GetRoleQuery) (uint, error) {
	var total int64
	var province model.Role
	statement := repo.db.WithContext(ctx).Model(&province)

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

	if query.Type != "" {
		statement = statement.Where("type = ?", query.Type)
	}

	if query.Platform != "" {
		statement = statement.Where("platform = ?", query.Platform)
	}

	if query.IsActiveFlag != "" {
		statement = statement.Where("is_active_flag = ?", query.IsActiveFlag)
	}

	err := statement.Count(&total).Error

	if err != nil {
		logger.Error(ctx, "Error GetRoleTotal", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "role", "repo"},
		})

		return uint(total), err
	}

	return uint(total), nil
}

func (repo *Repository) UpdateRole(ctx context.Context, role model.Role) (model.Role, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Save(&role).Error
	if err != nil {
		logger.Error(ctx, "Error UpdateRole", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "role", "repo"},
		})

		return role, err
	}

	return role, nil
}

func (repo *Repository) GetRole(ctx context.Context, query request.GetRoleQuery) (model.Role, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var role model.Role

	statement := repo.db.WithContext(ctx).Model(&role)

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

	if query.Platform != "" {
		statement = statement.Where("platform = ?", query.Platform)
	}

	if query.IsActiveFlag != "" {
		statement = statement.Where("is_active_flag = ?", query.IsActiveFlag)
	}

	if len(query.Preloads) > 0 {
		for _, p := range query.Preloads {
			statement = statement.Preload(p)
		}
	}

	err := tx.First(&role).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return role, lib.ErrorNotFound
	}

	if err != nil {
		logger.Error(ctx, "Error GetRole", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "role", "repo"},
		})

		return role, err
	}

	return role, nil
}

func (repo *Repository) DeleteRole(ctx context.Context, ID uint) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var role model.Role
	err := tx.WithContext(ctx).Delete(&role, ID).Error
	if err != nil {
		logger.Error(ctx, "Error DeleteRole", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "role", "repo"},
		})

		return err
	}

	return nil
}
