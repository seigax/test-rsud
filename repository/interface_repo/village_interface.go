package interface_repo

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

type VillageRepository interface {
	CreateVillage(ctx context.Context, village model.Village) (model.Village, error)
	GetVillages(ctx context.Context, query request.GetVillageQuery) ([]model.Village, error)
	GetVillageTotal(ctx context.Context, query request.GetVillageQuery) (uint, error)
	UpdateVillage(ctx context.Context, village model.Village) (model.Village, error)
	GetVillageByID(ctx context.Context, ID uint) (model.Village, error)
	GetVillagesByDistrictID(ctx context.Context, districtID uint) ([]model.Village, error)
	DeleteVillage(ctx context.Context, ID uint) error
}
