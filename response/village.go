package response

import "gitlab.com/erloom.id/libraries/go/backend-skeleton/model"

type GetVillageResponse struct {
	BasePaginateResponse
	Villages []model.Village `json:"villages"`
}
