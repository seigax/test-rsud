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

func (repo *Repository) CreateMenu(ctx context.Context, menu model.Menu) (model.Menu, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Create(&menu).Error
	if err != nil {
		logger.Error(ctx, "Error CreateMenu", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "menu", "repo"},
		})

		return menu, err
	}

	return menu, nil
}

func (repo *Repository) GetMenus(ctx context.Context, query request.GetMenuQuery) ([]model.Menu, error) {

	var res []model.Menu
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

	if query.ParentMenuID != 0 {
		statement = statement.Where("parent_menu_id = ?", query.ParentMenuID)
	}

	if query.Level != "" {
		statement = statement.Where("level = ?", query.Level)
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
		logger.Error(ctx, "Error GetMenu", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "menu", "repo"},
		})

		return res, err
	}

	return res, nil
}

func (repo *Repository) GetMenuTotal(ctx context.Context, query request.GetMenuQuery) (uint, error) {

	var total int64
	var menu model.Menu
	statement := repo.db.WithContext(ctx).Model(&menu)

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

	if query.ParentMenuID != 0 {
		statement = statement.Where("parent_menu_id = ?", query.ParentMenuID)
	}

	if query.Level != "" {
		statement = statement.Where("level = ?", query.Code)
	}

	if query.IsActiveFlag != "" {
		statement = statement.Where("is_active_flag = ?", query.IsActiveFlag)
	}

	err := statement.Count(&total).Error

	if err != nil {
		logger.Error(ctx, "Error GetMenuTotal", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "menu", "repo"},
		})

		return uint(total), err
	}

	return uint(total), nil
}

func (repo *Repository) UpdateMenu(ctx context.Context, menu model.Menu) (model.Menu, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Save(&menu).Error
	if err != nil {
		logger.Error(ctx, "Error UpdateMenu", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "menu", "repo"},
		})

		return menu, err
	}

	return menu, nil
}

func (repo *Repository) GetMenu(ctx context.Context, query request.GetMenuQuery) (menu model.Menu, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	statement := tx.WithContext(ctx).Model(&menu)

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

	if query.ParentMenuID != 0 {
		statement = statement.Where("parent_menu_id = ?", query.ParentMenuID)
	}

	if query.Level != "" {
		statement = statement.Where("level = ?", query.Level)
	}

	if query.IsActiveFlag != "" {
		statement = statement.Where("is_active_flag = ?", query.IsActiveFlag)
	}

	if len(query.Preloads) > 0 {
		for _, p := range query.Preloads {
			statement = statement.Preload(p)
		}
	}

	err = statement.First(&menu).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return menu, lib.ErrorNotFound
	}

	if err != nil {
		logger.Error(ctx, "Error GetMenu", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "menu", "repo"},
		})

		return menu, err
	}

	return menu, nil
}

func (repo *Repository) DeleteMenu(ctx context.Context, ID uint) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var menu model.Menu
	err := tx.WithContext(ctx).Delete(&menu, ID).Error
	if err != nil {
		logger.Error(ctx, "Error DeleteMenu", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "menu", "repo"},
		})

		return err
	}

	return nil
}
