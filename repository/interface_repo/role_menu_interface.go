package interface_repo

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
)

type RoleMenuRepository interface {
	CreateRoleMenu(ctx context.Context, roleMenu model.RoleMenu) (model.RoleMenu, error)
	GetRoleMenusByRoleID(ctx context.Context, roleID uint) ([]model.RoleMenu, error)
	GetRoleMenuByRoleIDAndMenuID(ctx context.Context, roleID uint, menuID uint) (model.RoleMenu, error)
	GetRoleMenusWithRoleUrlByRoleID(ctx context.Context, roleID uint) ([]model.RoleMenuWithRoleUrl, error)
	DeleteRoleMenu(ctx context.Context, ID uint) error
	DeleteAllRoleMenu(ctx context.Context, roleID uint) error
}
