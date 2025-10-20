package response

import "gitlab.com/erloom.id/libraries/go/backend-skeleton/model"

type GetSystemParameterResponse struct {
	BasePaginateResponse
	SystemParameters []model.SystemParameter `json:"system_parameters"`
}
