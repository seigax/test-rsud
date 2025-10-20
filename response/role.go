package response

import "gitlab.com/erloom.id/libraries/go/backend-skeleton/model"

type GetRoleResponse struct {
	BasePaginateResponse
	Roles []model.Role `json:"roles"`
}
