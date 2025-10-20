package repository

import (
	"context"
	"errors"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/logger"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gorm.io/gorm"
)

func (repo *Repository) CreateUserRole(ctx context.Context, userrole model.UserRole) (model.UserRole, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Create(&userrole).Error
	if err != nil {
		logger.Error(ctx, "Error CreateUserRole", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "user_role", "repo"},
		})

		return userrole, err
	}

	return userrole, nil
}

func (repo *Repository) UpdateUserRole(ctx context.Context, userrole model.UserRole) (model.UserRole, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Save(&userrole).Error
	if err != nil {
		logger.Error(ctx, "Error UpdateUserRole", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "user_role", "repo"},
		})

		return userrole, err
	}

	return userrole, nil
}

func (repo *Repository) GetUserRoleByID(ctx context.Context, ID uint) (model.UserRole, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var userRole model.UserRole

	err := tx.WithContext(ctx).Where("id = ?", ID).First(&userRole).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return userRole, lib.ErrorNotFound
	}

	if err != nil {
		logger.Error(ctx, "Error GetUserRoleByID", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "user_role", "repo"},
		})

		return userRole, err
	}

	return userRole, nil
}

func (repo *Repository) GetUserRoleByUserID(ctx context.Context, userID uint) (model.UserRole, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var userRole model.UserRole

	err := tx.WithContext(ctx).Where("user_id = ?", userID).First(&userRole).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return userRole, lib.ErrorNotFound
	}

	if err != nil {
		logger.Error(ctx, "Error GetUserRoleByUserID", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "user_role", "repo"},
		})

		return userRole, err
	}

	return userRole, nil
}

func (repo *Repository) GetUserRolesByUserID(ctx context.Context, userID uint) ([]model.UserRole, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var userRole []model.UserRole

	err := tx.WithContext(ctx).Where("user_id = ?", userID).Find(&userRole).Error
	if err != nil {
		logger.Error(ctx, "Error GetUserRoleByUserID", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "user_role", "repo"},
		})

		return userRole, err
	}

	return userRole, nil
}

func (repo *Repository) GetUserRolesWithRoleDetailByUserID(ctx context.Context, userID uint) ([]model.UserRoleWithRoleDetail, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var userRole []model.UserRoleWithRoleDetail

	err := tx.WithContext(ctx).
		Table("user_role as ur").
		Select([]string{
			"ur.id",
			"ur.user_id",
			"ur.role_id",
			"ur.created_at",
			"ur.created_by",
			"ur.updated_at",
			"ur.updated_by",
			"r.name as role_name",
			"r.type as role_type",
			"r.platform as role_platform",
		}).
		Joins("JOIN role as r on r.id = ur.role_id").
		Where("ur.user_id = ?", userID).
		Find(&userRole).Error

	if err != nil {
		logger.Error(ctx, "Error GetUserRolesWithRoleNameByUserID", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "user_role", "repo"},
		})

		return userRole, err
	}

	return userRole, nil
}

func (repo *Repository) DeleteUserRole(ctx context.Context, ID uint) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var userRole model.UserRole
	err := tx.WithContext(ctx).Delete(&userRole, ID).Error
	if err != nil {
		logger.Error(ctx, "Error DeleteUserRole", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "user_role", "repo"},
		})

		return err
	}

	return nil
}

func (repo *Repository) DeleteUserRoleByUserID(ctx context.Context, userID uint) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var userRole model.UserRole
	err := tx.WithContext(ctx).Where("user_id = ?", userID).Delete(&userRole).Error
	if err != nil {
		logger.Error(ctx, "Error DeleteUserRoleByUserID", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "user_role", "repo"},
		})

		return err
	}

	return nil
}
