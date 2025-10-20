package handler

import (
	"encoding/json"
	"net/http"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/validation"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

func (handler *Handler) GetFarmerPhone(rw http.ResponseWriter, r *http.Request) {
	var payload request.PhoneRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		WriteError(rw, lib.ErrorInvalidParameter)
		return
	}

	err = validation.Validator.Struct(payload)
	if err != nil {
		WriteError(rw, lib.ErrorInvalidParameter)
		return
	}

	resp, err := handler.BackendSkeleton.Usecase.GetFarmerByPhone(r.Context(), payload)
	if err != nil {
		WriteError(rw, err)
		return
	}

	WriteSuccess(rw, resp, "Success Send OTP!", ResponseMeta{HTTPStatus: http.StatusOK})
}
