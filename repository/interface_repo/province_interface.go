package interface_repo

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

type ProvinceRepository interface {
	CreateProvince(ctx context.Context, province model.Province) (model.Province, error)
	GetProvinces(ctx context.Context, query request.GetProvinceQuery) ([]model.Province, error)
	GetProvinceTotal(ctx context.Context, query request.GetProvinceQuery) (uint, error)
	UpdateProvince(ctx context.Context, province model.Province) (model.Province, error)
	UpdateProvinceAddTotalCity(ctx context.Context, provinceID uint, add int) error
	UpdateProvinceAddTotalDistrict(ctx context.Context, provinceID uint, add int) error
	UpdateProvinceAddTotalVillage(ctx context.Context, provinceID uint, add int) error
	GetProvince(ctx context.Context, query request.GetProvinceQuery) (province model.Province, err error)
	DeleteProvince(ctx context.Context, ID uint) error
}
