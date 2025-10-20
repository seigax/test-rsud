package response

import "gitlab.com/erloom.id/libraries/go/backend-skeleton/model"

type GetDistrictResponse struct {
	BasePaginateResponse
	Districts []model.District `json:"districts"`
}

type GetDistrictDetailWithTreeResponse struct {
	model.District
	Village []model.Village `json:"villages"`
}
