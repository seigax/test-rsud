package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/validation"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/response"
)

func (handler *Handler) GetMenusWithChild(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	var roleID int

	roleID, _ = strconv.Atoi(queryParams.Get("role_id"))

	res, err := handler.BackendSkeleton.Usecase.GetMenusWithChild(r.Context(), uint(roleID))
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, res, "success", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) SaveRoleMenu(rw http.ResponseWriter, r *http.Request) {
	var payload request.CreateRoleMenuRequest

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

	currentUserID := r.Context().Value("CurrentUserID").(uint)
	payload.CreatedBy = currentUserID

	err = handler.BackendSkeleton.Usecase.SaveRoleMenu(r.Context(), payload)
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, response.EmptyResponse{}, "Create Role Menu Success", ResponseMeta{HTTPStatus: http.StatusOK})
}
