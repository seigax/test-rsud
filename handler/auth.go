package handler

import (
	"encoding/json"
	"net/http"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/validation"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

// func (handler *Handler) RegisterFarmer(rw http.ResponseWriter, r *http.Request) {
// 	var payload request.RegisterFarmerRequest

// 	err := json.NewDecoder(r.Body).Decode(&payload)
// 	if err != nil {
// 		handler.WriteError(r.Context(), rw, lib.ErrorInvalidParameter)
// 		return
// 	}

// 	err = validation.Validator.Struct(payload)
// 	if err != nil {
// 		handler.WriteError(r.Context(), rw, err)
// 		return
// 	}

// 	_, err = handler.BackendSkeleton.Usecase.RegisterFarmer(r.Context(), payload)
// 	if err != nil {
// 		WriteError(rw, err)
// 		return
// 	}

// 	WriteSuccess(rw, response.EmptyResponse{}, "Register Success! Please Check Your Email And Validate Your Account!", ResponseMeta{HTTPStatus: http.StatusOK})
// }

func (handler *Handler) Login(rw http.ResponseWriter, r *http.Request) {
	var payload request.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		handler.WriteError(r.Context(), rw, lib.ErrorInvalidParameter)
		return
	}

	err = validation.Validator.Struct(payload)
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	resp, err := handler.BackendSkeleton.Usecase.Login(r.Context(), payload)
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, resp, "Login Success!", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) Logout(rw http.ResponseWriter, r *http.Request) {
	session := r.Context().Value("CurrentSession").(model.Session)
	err := handler.BackendSkeleton.Usecase.ExpireToken(r.Context(), session)
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, nil, "Logout Success!", ResponseMeta{HTTPStatus: http.StatusOK})
}
