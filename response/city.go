package response

import "gitlab.com/erloom.id/libraries/go/backend-skeleton/model"

type GetCityResponse struct {
	BasePaginateResponse
	Citys []model.City `json:"citys"`
}

type GetCityDetailWithTreeResponse struct {
	model.City
	District []GetDistrictDetailWithTreeResponse `json:"districts"`
}
