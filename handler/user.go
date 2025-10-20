package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/validation"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

func (handler *Handler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	var payload request.CreateUserRequest

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

	res, err := handler.BackendSkeleton.Usecase.CreateUser(r.Context(), payload)
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, res, "success", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	var limit, page int

	limit, _ = strconv.Atoi(queryParams.Get("limit"))
	page, _ = strconv.Atoi(queryParams.Get("page"))

	sort := strings.Split(queryParams.Get("sort"), ",")

	var query request.GetUserQuery

	query.Limit = uint(limit)
	query.Page = uint(page)
	query.Search = queryParams.Get("search")
	query.Sort = sort

	res, err := handler.BackendSkeleton.Usecase.GetUsers(r.Context(), query)
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	resp := ResponseMeta{HTTPStatus: http.StatusOK}
	resp.SerializeFromResponse(res.BasePaginateResponse)

	WriteSuccess(rw, res, "success", resp)
}

func (handler *Handler) GetUserDetail(rw http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "ID")
	id, _ := strconv.ParseUint(idString, 10, 32)

	res, err := handler.BackendSkeleton.Usecase.GetUserDetail(r.Context(), uint(id))
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, res, "success", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	var payload request.UpdateUserRequest

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

	idString := chi.URLParam(r, "ID")
	id, _ := strconv.ParseUint(idString, 10, 32)

	payload.UserID = uint(id)

	currentUserID := r.Context().Value("CurrentUserID").(uint)
	payload.UpdatedBy = currentUserID

	res, err := handler.BackendSkeleton.Usecase.UpdateUser(r.Context(), payload)
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, res, "success", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "ID")
	id, _ := strconv.ParseUint(idString, 10, 32)

	err := handler.BackendSkeleton.Usecase.DeleteUser(r.Context(), uint(id))
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, nil, "success", ResponseMeta{HTTPStatus: http.StatusOK})
}
