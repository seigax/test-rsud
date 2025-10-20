package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/validation"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

func (handler *Handler) GetPhonesFarmer(rw http.ResponseWriter, r *http.Request) {

	userID, ok := r.Context().Value("CurrentUserID").(uint)
	if !ok {
		WriteError(rw, lib.ErrorInternalServer)
		return
	}

	resp, err := handler.BackendSkeleton.Usecase.GetListUserPhone(r.Context(), userID)
	if err != nil {
		WriteError(rw, err)
		return
	}

	WriteSuccess(rw, resp, "Success Get Phone Farmer!", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) AddPhonesFarmer(rw http.ResponseWriter, r *http.Request) {
	var payload request.ReqAddPhone

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

	resp, err := handler.BackendSkeleton.Usecase.AddUserPhone(r.Context(), payload)
	if err != nil {
		WriteError(rw, err)
		return
	}

	WriteSuccess(rw, resp, "Success Add Phone Farmer!", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) DeletePhone(rw http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "ID")
	id, _ := strconv.ParseUint(idString, 10, 32)
	if id == 0 {
		WriteError(rw, lib.ErrorInvalidParameter)
		return
	}

	err := handler.BackendSkeleton.Usecase.DeleteUserPhone(r.Context(), uint(id))
	if err != nil {
		WriteError(rw, err)
		return
	}
	WriteSuccess(rw, nil, "Success Delete Phone ", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) UpdatePhonesFarmer(rw http.ResponseWriter, r *http.Request) {
	var payload request.ReqUpdatePhone

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

	idString := chi.URLParam(r, "ID")
	id, _ := strconv.ParseUint(idString, 10, 32)
	if id == 0 {
		WriteError(rw, lib.ErrorInvalidParameter)
		return
	}

	payload.Id = uint(id)

	resp, err := handler.BackendSkeleton.Usecase.UpdateUserPhone(r.Context(), payload)
	if err != nil {
		WriteError(rw, err)
		return
	}

	WriteSuccess(rw, resp, "Success Update Phone Farmer!", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) ChangeActivePhoneFarmer(rw http.ResponseWriter, r *http.Request) {
	var payload request.ReqUpdatePhone

	userID, ok := r.Context().Value("CurrentUserID").(uint)
	if !ok {
		WriteError(rw, lib.ErrorInternalServer)
		return
	}
	payload.UserId = userID

	idString := chi.URLParam(r, "ID")
	id, _ := strconv.ParseUint(idString, 10, 32)
	if id == 0 {
		WriteError(rw, lib.ErrorInvalidParameter)
		return
	}

	payload.Id = uint(id)

	err := handler.BackendSkeleton.Usecase.ChangeActiveUserPhone(r.Context(), payload)
	if err != nil {
		WriteError(rw, err)
		return
	}

	WriteSuccess(rw, nil, "Success Change Active Phone Farmer!", ResponseMeta{HTTPStatus: http.StatusOK})
}
