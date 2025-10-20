package interface_repo

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

type MenuRepository interface {
	CreateMenu(ctx context.Context, menu model.Menu) (model.Menu, error)
	GetMenus(ctx context.Context, query request.GetMenuQuery) ([]model.Menu, error)
	GetMenuTotal(ctx context.Context, query request.GetMenuQuery) (uint, error)
	UpdateMenu(ctx context.Context, menu model.Menu) (model.Menu, error)
	GetMenu(ctx context.Context, query request.GetMenuQuery) (model.Menu, error)
	DeleteMenu(ctx context.Context, ID uint) error
}
