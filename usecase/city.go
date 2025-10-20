package usecase

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/response"
)

func (usecase *Usecase) CreateCity(ctx context.Context, payload request.CreateCityRequest) (city model.City, err error) {
	err = usecase.repo.Transaction(ctx, func(ctx context.Context) error {
		city = model.City{
			ProvinceID:   payload.ProvinceID,
			Name:         payload.Name,
			IsActiveFlag: "Y",
			CreatedAt:    time.Now(),
			CreatedBy:    payload.CreatedBy,
			UpdatedAt:    time.Now(),
			UpdatedBy:    payload.CreatedBy,
		}

		city, err = usecase.repo.CreateCity(ctx, city)
		if err != nil {
			return err
		}

		city.Code = fmt.Sprintf("CITY-%06d", city.ID)

		city, err = usecase.repo.UpdateCity(ctx, city)
		if err != nil {
			return err
		}

		err = usecase.repo.UpdateProvinceAddTotalCity(ctx, city.ProvinceID, 1)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

func (usecase *Usecase) GetCitys(ctx context.Context, query request.GetCityQuery) (response.GetCityResponse, error) {
	var res response.GetCityResponse
	citys, err := usecase.repo.GetCitys(ctx, query)
	if err != nil {
		return res, err
	}

	total, err := usecase.repo.GetCityTotal(ctx, query)
	if err != nil {
		return res, err
	}

	res.Citys = citys
	res.Total = total
	res.Limit = query.Limit
	res.Offset = query.GetOffset()
	res.Page = query.Page

	return res, nil
}

func (usecase *Usecase) GetCityDetail(ctx context.Context, ID uint) (model.City, error) {
	city, err := usecase.repo.GetCity(ctx, request.GetCityQuery{
		City: model.City{ID: ID},
	})
	if err != nil {
		return city, err
	}

	return city, nil
}

func (usecase *Usecase) UpdateCity(ctx context.Context, payload request.UpdateCityRequest) (model.City, error) {
	city, err := usecase.repo.GetCity(ctx, request.GetCityQuery{
		City: model.City{ID: payload.ID},
	})
	if err != nil {
		return city, err
	}

	city.Name = payload.Name
	city.UpdatedAt = time.Now()
	city.UpdatedBy = payload.UpdatedBy

	res, err := usecase.repo.UpdateCity(ctx, city)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (usecase *Usecase) DeleteCity(ctx context.Context, ID uint) error {
	err := usecase.repo.Transaction(ctx, func(ctx context.Context) error {
		city, err := usecase.repo.GetCity(ctx, request.GetCityQuery{
			City: model.City{ID: ID},
		})
		if err != nil {
			return err
		}

		err = usecase.repo.DeleteCity(ctx, ID)
		if err != nil {
			return err
		}

		err = usecase.repo.UpdateProvinceAddTotalCity(ctx, city.ProvinceID, -1)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}
