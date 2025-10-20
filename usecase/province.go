package usecase

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/response"
)

func (usecase *Usecase) CreateProvince(ctx context.Context, payload request.CreateProvinceRequest) (province model.Province, err error) {
	err = usecase.repo.Transaction(ctx, func(ctx context.Context) error {
		province = model.Province{
			Name:         payload.Name,
			IsActiveFlag: "Y",
			CreatedAt:    time.Now(),
			CreatedBy:    payload.CreatedBy,
			UpdatedAt:    time.Now(),
			UpdatedBy:    payload.CreatedBy,
		}

		province, err = usecase.repo.CreateProvince(ctx, province)
		if err != nil {
			return err
		}

		province.Code = fmt.Sprintf("PROV-%04d", province.ID)

		province, err = usecase.repo.UpdateProvince(ctx, province)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

func (usecase *Usecase) GetProvinces(ctx context.Context, query request.GetProvinceQuery) (response.GetProvinceResponse, error) {
	var res response.GetProvinceResponse

	provinces, err := usecase.repo.GetProvinces(ctx, query)
	if err != nil {
		return res, err
	}

	total, err := usecase.repo.GetProvinceTotal(ctx, query)
	if err != nil {
		return res, err
	}

	res.Provinces = provinces
	res.Total = total
	res.Limit = query.Limit
	res.Offset = query.GetOffset()
	res.Page = query.Page

	return res, nil
}

func (usecase *Usecase) GetProvinceDetail(ctx context.Context, ID uint) (model.Province, error) {
	province, err := usecase.repo.GetProvince(ctx, request.GetProvinceQuery{Province: model.Province{ID: ID}})
	if err != nil {
		return province, err
	}

	return province, nil
}

func (usecase *Usecase) GetProvinceDetailWithTree(ctx context.Context, ID uint) (resp response.GetProvinceDetailWithTreeResponse, err error) {
	province, err := usecase.repo.GetProvince(ctx, request.GetProvinceQuery{Province: model.Province{ID: ID}})
	if err != nil {
		return
	}

	resp.Province = province

	citys, _ := usecase.repo.GetCitys(ctx, request.GetCityQuery{
		City: model.City{ProvinceID: province.ID},
	})

	for _, city := range citys {
		respCity := response.GetCityDetailWithTreeResponse{
			City: city,
		}

		districts, _ := usecase.repo.GetDistricts(ctx, request.GetDistrictQuery{District: model.District{CityID: city.ID}})

		for _, district := range districts {
			respDistrict := response.GetDistrictDetailWithTreeResponse{
				District: district,
			}

			villages, _ := usecase.repo.GetVillagesByDistrictID(ctx, district.ID)
			respDistrict.Village = villages

			respCity.District = append(respCity.District, respDistrict)
		}

		resp.City = append(resp.City, respCity)
	}

	return
}

func (usecase *Usecase) UpdateProvince(ctx context.Context, payload request.UpdateProvinceRequest) (model.Province, error) {
	province, err := usecase.repo.GetProvince(ctx, request.GetProvinceQuery{Province: model.Province{ID: payload.ID}})
	if err != nil {
		return province, err
	}

	province.Name = payload.Name
	province.UpdatedAt = time.Now()
	province.UpdatedBy = payload.UpdatedBy

	res, err := usecase.repo.UpdateProvince(ctx, province)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (usecase *Usecase) DeleteProvince(ctx context.Context, ID uint) error {
	err := usecase.repo.DeleteProvince(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
