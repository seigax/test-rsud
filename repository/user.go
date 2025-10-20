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

func (repo *Repository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Create(&user).Error
	if err != nil {
		logger.Error(ctx, "Error CreateUser", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "user", "repo"},
		})

		return user, err
	}

	return user, nil
}

func (repo *Repository) GetUsers(ctx context.Context, query request.GetUserQuery) ([]model.User, error) {

	var res []model.User
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
		logger.Error(ctx, "Error GetUser", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "user", "repo"},
		})

		return res, err
	}

	return res, nil
}

func (repo *Repository) GetUserTotal(ctx context.Context, query request.GetUserQuery) (uint, error) {

	var total int64
	var user model.User
	statement := repo.db.WithContext(ctx).Model(&user)

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
		logger.Error(ctx, "Error GetUserTotal", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "user", "repo"},
		})

		return uint(total), err
	}

	return uint(total), nil
}

func (repo *Repository) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err := tx.WithContext(ctx).Save(&user).Error
	if err != nil {
		logger.Error(ctx, "Error UpdateUser", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "user", "repo"},
		})

		return user, err
	}

	return user, nil
}

func (repo *Repository) GetUserByID(ctx context.Context, ID uint) (user model.User, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err = tx.WithContext(ctx).First(&user, ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, lib.ErrorNotFound
	}
	if err != nil {
		logger.Error(ctx, "Error GetUserByID", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "user", "repo"},
		})

		return user, err
	}

	return user, nil
}

func (repo *Repository) GetUserByEmail(ctx context.Context, email string) (user model.User, err error) {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	err = tx.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, lib.ErrorNotFound
	}

	if err != nil {
		logger.Error(ctx, "Error GetUserByEmail", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "user", "repo"},
		})

		return user, err
	}

	return user, nil
}

func (repo *Repository) DeleteUser(ctx context.Context, ID uint) error {
	tx, ok := ctx.Value("Trx").(*lib.Database)
	if !ok {
		tx = repo.db
	}

	var user model.User
	err := tx.WithContext(ctx).Delete(&user, ID).Error
	if err != nil {
		logger.Error(ctx, "Error DeleteUser", map[string]interface{}{
			"error": err,
			"tags":  []string{"postgres", "user", "repo"},
		})

		return err
	}

	return nil
}
