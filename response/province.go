package response

import "gitlab.com/erloom.id/libraries/go/backend-skeleton/model"

type GetProvinceResponse struct {
	BasePaginateResponse
	Provinces []model.Province `json:"provinces"`
}

type GetProvinceDetailWithTreeResponse struct {
	model.Province
	City []GetCityDetailWithTreeResponse `json:"citys"`
}
