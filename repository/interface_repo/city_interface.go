package interface_repo

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

type CityRepository interface {
	CreateCity(ctx context.Context, city model.City) (model.City, error)
	GetCitys(ctx context.Context, query request.GetCityQuery) ([]model.City, error)
	GetCityTotal(ctx context.Context, query request.GetCityQuery) (uint, error)
	UpdateCity(ctx context.Context, city model.City) (model.City, error)
	GetCity(ctx context.Context, query request.GetCityQuery) (model.City, error)
	DeleteCity(ctx context.Context, ID uint) error
}
