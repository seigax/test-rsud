package interface_repo

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
)

type UserRoleRepository interface {
	CreateUserRole(ctx context.Context, userrole model.UserRole) (model.UserRole, error)
	UpdateUserRole(ctx context.Context, userrole model.UserRole) (model.UserRole, error)
	GetUserRoleByID(ctx context.Context, ID uint) (model.UserRole, error)
	GetUserRoleByUserID(ctx context.Context, userID uint) (model.UserRole, error)
	GetUserRolesWithRoleDetailByUserID(ctx context.Context, userID uint) ([]model.UserRoleWithRoleDetail, error)
	GetUserRolesByUserID(ctx context.Context, userID uint) ([]model.UserRole, error)
	DeleteUserRole(ctx context.Context, ID uint) error
	DeleteUserRoleByUserID(ctx context.Context, userID uint) error
}
