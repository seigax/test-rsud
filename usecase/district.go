package usecase

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/response"
)

func (usecase *Usecase) CreateDistrict(ctx context.Context, payload request.CreateDistrictRequest) (district model.District, err error) {
	err = usecase.repo.Transaction(ctx, func(ctx context.Context) error {
		district = model.District{
			CityID:       payload.CityID,
			Name:         payload.Name,
			IsActiveFlag: "Y",
			CreatedAt:    time.Now(),
			CreatedBy:    payload.CreatedBy,
			UpdatedAt:    time.Now(),
			UpdatedBy:    payload.CreatedBy,
		}

		district, err = usecase.repo.CreateDistrict(ctx, district)
		if err != nil {
			return err
		}

		district.Code = fmt.Sprintf("DISTRICT-%08d", district.ID)

		district, err = usecase.repo.UpdateDistrict(ctx, district)
		if err != nil {
			return err
		}

		city, err := usecase.repo.GetCity(ctx, request.GetCityQuery{
			City: model.City{ID: district.CityID},
		})
		if err != nil {
			return err
		}

		err = usecase.repo.UpdateProvinceAddTotalDistrict(ctx, city.ProvinceID, 1)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

func (usecase *Usecase) GetDistricts(ctx context.Context, query request.GetDistrictQuery) (response.GetDistrictResponse, error) {
	var res response.GetDistrictResponse

	districts, err := usecase.repo.GetDistricts(ctx, query)
	if err != nil {
		return res, err
	}

	total, err := usecase.repo.GetDistrictTotal(ctx, query)
	if err != nil {
		return res, err
	}

	res.Districts = districts
	res.Total = total
	res.Limit = query.Limit
	res.Offset = query.GetOffset()
	res.Page = query.Page

	return res, nil
}

func (usecase *Usecase) GetDistrictDetail(ctx context.Context, ID uint) (model.District, error) {
	district, err := usecase.repo.GetDistrict(ctx, request.GetDistrictQuery{District: model.District{ID: ID}})
	if err != nil {
		return district, err
	}

	return district, nil
}

func (usecase *Usecase) UpdateDistrict(ctx context.Context, payload request.UpdateDistrictRequest) (model.District, error) {
	district, err := usecase.repo.GetDistrict(ctx, request.GetDistrictQuery{District: model.District{ID: payload.ID}})
	if err != nil {
		return district, err
	}

	district.Name = payload.Name
	district.UpdatedAt = time.Now()
	district.UpdatedBy = payload.UpdatedBy

	res, err := usecase.repo.UpdateDistrict(ctx, district)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (usecase *Usecase) DeleteDistrict(ctx context.Context, ID uint) error {
	err := usecase.repo.Transaction(ctx, func(ctx context.Context) error {
		district, err := usecase.repo.GetDistrict(ctx, request.GetDistrictQuery{
			District: model.District{ID: ID},
		})
		if err != nil {
			return err
		}

		err = usecase.repo.DeleteDistrict(ctx, ID)
		if err != nil {
			return err
		}

		city, err := usecase.repo.GetCity(ctx, request.GetCityQuery{
			City: model.City{ID: district.CityID},
		})
		if err != nil {
			return err
		}

		err = usecase.repo.UpdateProvinceAddTotalDistrict(ctx, city.ProvinceID, -1)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}
