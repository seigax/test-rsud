package response

import "gitlab.com/erloom.id/libraries/go/backend-skeleton/model"

type GetErrorMessageResponse struct {
	BasePaginateResponse
	ErrorMessages []model.ErrorMessage `json:"error_messages"`
}
