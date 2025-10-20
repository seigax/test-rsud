package handler

import (
	"encoding/json"
	"net/http"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/validation"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

func (handler *Handler) CheckOTPLoginFarmer(rw http.ResponseWriter, r *http.Request) {
	var payload request.OTPRequest

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

	resp, err := handler.BackendSkeleton.Usecase.CheckOTPLogin(r.Context(), payload)
	if err != nil {
		WriteError(rw, err)
		return
	}

	WriteSuccess(rw, resp, "Check OTP Success!", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) ChangePassword(rw http.ResponseWriter, r *http.Request) {
	var payload request.ChangePassword

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

	userID, ok := r.Context().Value("CurrentUserID").(uint)
	if !ok {
		WriteError(rw, lib.ErrorInternalServer)
		return
	}

	payload.UserId = userID

	err = handler.BackendSkeleton.Usecase.ChangePassword(r.Context(), payload)
	if err != nil {
		WriteError(rw, err)
		return
	}

	WriteSuccess(rw, nil, "Success Change Password!", ResponseMeta{HTTPStatus: http.StatusOK})
}
