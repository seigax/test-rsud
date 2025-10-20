package interface_repo

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

type SystemParameterRepository interface {
	CreateSystemParameter(ctx context.Context, systemParameter model.SystemParameter) (model.SystemParameter, error)
	GetSystemParameters(ctx context.Context, query request.GetSystemParameterQuery) ([]model.SystemParameter, error)
	GetSystemParameterTotal(ctx context.Context, query request.GetSystemParameterQuery) (uint, error)
	UpdateSystemParameter(ctx context.Context, systemParameter model.SystemParameter) (model.SystemParameter, error)
	GetSystemParameter(ctx context.Context, query request.GetSystemParameterQuery) (model.SystemParameter, error)
	DeleteSystemParameter(ctx context.Context, ID uint) error
}
