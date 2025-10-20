package repository

import (
	"context"
	"errors"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/logger"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gorm.io/gorm"
)

func (repo *Repository) CreateRoleMenu(ctx context.Context, roleMenu model.RoleMenu) (model.RoleMenu, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Create(&roleMenu).Error
	if err != nil {
		logger.Error(ctx, "Error CreateRoleMenu", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "roleMenu", "repo"},
		})

		return roleMenu, err
	}

	return roleMenu, nil
}

func (repo *Repository) GetRoleMenusByRoleID(ctx context.Context, roleID uint) (roleMenu []model.RoleMenu, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err = tx.WithContext(ctx).
		Where("role_id = ?", roleID).
		Find(&roleMenu).Error
	if err != nil {
		logger.Error(ctx, "Error GetRoleMenuByID", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "RoleMenu", "repo"},
		})

		return
	}

	return
}

func (repo *Repository) GetRoleMenuByRoleIDAndMenuID(ctx context.Context, roleID uint, menuID uint) (roleMenu model.RoleMenu, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err = tx.WithContext(ctx).Where("role_id = ? AND menu_id = ?", roleID, menuID).First(&roleMenu).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return roleMenu, lib.ErrorNotFound
	}

	if err != nil {
		logger.Error(ctx, "Error GetRoleMenuByRoleIDAndMenuID", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "RoleMenu", "repo"},
		})

		return roleMenu, err
	}

	return roleMenu, nil
}

func (repo *Repository) GetRoleMenusWithRoleUrlByRoleID(ctx context.Context, roleID uint) (roleMenu []model.RoleMenuWithRoleUrl, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err = tx.WithContext(ctx).
		Table("role_menu as rm").
		Select([]string{
			"rm.id",
			"rm.role_id",
			"rm.menu_id",
			"rm.created_at",
			"rm.created_by",
			"rm.updated_at",
			"rm.updated_by",
			"m.url as menu_url",
		}).
		Joins("JOIN menu as m on m.id = rm.menu_id").
		Where("rm.role_id = ?", roleID).
		Find(&roleMenu).Error

	if err != nil {
		logger.Error(ctx, "Error GetRoleMenusWithRoleUrlByRoleID", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "RoleMenu", "repo"},
		})

		return
	}

	return
}

func (repo *Repository) DeleteRoleMenu(ctx context.Context, ID uint) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var roleMenu model.RoleMenu
	err := tx.WithContext(ctx).Delete(&roleMenu, ID).Error
	if err != nil {
		logger.Error(ctx, "Error DeleteRoleMenu", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "roleMenu", "repo"},
		})

		return err
	}

	return nil
}

func (repo *Repository) DeleteAllRoleMenu(ctx context.Context, roleID uint) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var roleMenu model.RoleMenu
	err := tx.WithContext(ctx).
		Where("role_id = ?", roleID).
		Delete(&roleMenu).Error
	if err != nil {
		logger.Error(ctx, "Error DeleteRoleMenu", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "roleMenu", "repo"},
		})

		return err
	}

	return nil
}
