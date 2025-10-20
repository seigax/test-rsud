package usecase

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/response"
)

func (usecase *Usecase) CreateVillage(ctx context.Context, payload request.CreateVillageRequest) (village model.Village, err error) {
	err = usecase.repo.Transaction(ctx, func(ctx context.Context) error {
		village = model.Village{
			DistrictID:   payload.DistrictID,
			PostalCode:   payload.PostalCode,
			Name:         payload.Name,
			IsActiveFlag: "Y",
			CreatedAt:    time.Now(),
			CreatedBy:    payload.CreatedBy,
			UpdatedAt:    time.Now(),
			UpdatedBy:    payload.CreatedBy,
		}

		village, err = usecase.repo.CreateVillage(ctx, village)
		if err != nil {
			return err
		}

		village.Code = fmt.Sprintf("VILL-%011d", village.ID)

		village, err = usecase.repo.UpdateVillage(ctx, village)
		if err != nil {
			return err
		}

		district, err := usecase.repo.GetDistrict(ctx, request.GetDistrictQuery{District: model.District{ID: village.DistrictID}})
		if err != nil {
			return err
		}

		city, err := usecase.repo.GetCity(ctx, request.GetCityQuery{
			City: model.City{ID: district.CityID},
		})
		if err != nil {
			return err
		}

		err = usecase.repo.UpdateProvinceAddTotalVillage(ctx, city.ProvinceID, 1)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

func (usecase *Usecase) GetVillages(ctx context.Context, query request.GetVillageQuery) (response.GetVillageResponse, error) {
	var res response.GetVillageResponse

	villages, err := usecase.repo.GetVillages(ctx, query)
	if err != nil {
		return res, err
	}

	total, err := usecase.repo.GetVillageTotal(ctx, query)
	if err != nil {
		return res, err
	}

	res.Villages = villages
	res.Total = total
	res.Limit = query.Limit
	res.Offset = query.GetOffset()
	res.Page = query.Page

	return res, nil
}

func (usecase *Usecase) GetVillageDetail(ctx context.Context, ID uint) (model.Village, error) {
	village, err := usecase.repo.GetVillageByID(ctx, ID)
	if err != nil {
		return village, err
	}

	return village, nil
}

func (usecase *Usecase) UpdateVillage(ctx context.Context, payload request.UpdateVillageRequest) (model.Village, error) {
	village, err := usecase.repo.GetVillageByID(ctx, payload.ID)
	if err != nil {
		return village, err
	}

	village.Name = payload.Name
	village.UpdatedAt = time.Now()
	village.UpdatedBy = payload.UpdatedBy

	res, err := usecase.repo.UpdateVillage(ctx, village)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (usecase *Usecase) DeleteVillage(ctx context.Context, ID uint) error {
	err := usecase.repo.Transaction(ctx, func(ctx context.Context) error {
		village, err := usecase.repo.GetVillageByID(ctx, ID)
		if err != nil {
			return err
		}

		err = usecase.repo.DeleteVillage(ctx, ID)
		if err != nil {
			return err
		}

		district, err := usecase.repo.GetDistrict(ctx, request.GetDistrictQuery{District: model.District{ID: village.DistrictID}})
		if err != nil {
			return err
		}

		city, err := usecase.repo.GetCity(ctx, request.GetCityQuery{
			City: model.City{ID: district.CityID},
		})
		if err != nil {
			return err
		}

		err = usecase.repo.UpdateProvinceAddTotalVillage(ctx, city.ProvinceID, -1)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}
