package interface_repo

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

type RoleRepository interface {
	CreateRole(ctx context.Context, role model.Role) (model.Role, error)
	GetRoles(ctx context.Context, query request.GetRoleQuery) ([]model.Role, error)
	GetRoleTotal(ctx context.Context, query request.GetRoleQuery) (uint, error)
	UpdateRole(ctx context.Context, role model.Role) (model.Role, error)
	GetRole(ctx context.Context, query request.GetRoleQuery) (model.Role, error)
	DeleteRole(ctx context.Context, ID uint) error
}
