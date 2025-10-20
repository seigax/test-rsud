package interface_repo

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

type DistrictRepository interface {
	CreateDistrict(ctx context.Context, district model.District) (model.District, error)
	GetDistricts(ctx context.Context, query request.GetDistrictQuery) ([]model.District, error)
	GetDistrictTotal(ctx context.Context, query request.GetDistrictQuery) (uint, error)
	UpdateDistrict(ctx context.Context, district model.District) (model.District, error)
	GetDistrict(ctx context.Context, query request.GetDistrictQuery) (model.District, error)
	DeleteDistrict(ctx context.Context, ID uint) error
}
