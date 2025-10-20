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
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/response"
)

func (handler *Handler) CreateMenu(rw http.ResponseWriter, r *http.Request) {
	var payload request.CreateMenuRequest

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

	res, err := handler.BackendSkeleton.Usecase.CreateMenu(r.Context(), payload)
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, res, "Create Menu Success", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) GetMenus(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	var limit, page int

	limit, _ = strconv.Atoi(queryParams.Get("limit"))
	page, _ = strconv.Atoi(queryParams.Get("page"))

	sort := strings.Split(queryParams.Get("sort"), ",")

	var query request.GetMenuQuery

	query.Limit = uint(limit)
	query.Page = uint(page)
	query.Search = queryParams.Get("search")
	query.Sort = sort

	res, err := handler.BackendSkeleton.Usecase.GetMenus(r.Context(), query)
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	resp := ResponseMeta{HTTPStatus: http.StatusOK}
	resp.SerializeFromResponse(res.BasePaginateResponse)

	WriteSuccess(rw, res, "success", resp)
}

func (handler *Handler) GetMenuDetail(rw http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "ID")
	id, _ := strconv.ParseUint(idString, 10, 32)

	res, err := handler.BackendSkeleton.Usecase.GetMenuDetail(r.Context(), uint(id))
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, res, "success", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) GetMenuDetailWithChild(rw http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "ID")
	id, _ := strconv.ParseUint(idString, 10, 32)

	res, err := handler.BackendSkeleton.Usecase.GetMenuDetailWithChild(r.Context(), uint(id))
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, res, "success", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) UpdateMenu(rw http.ResponseWriter, r *http.Request) {
	var payload request.UpdateMenuRequest

	idString := chi.URLParam(r, "ID")
	id, _ := strconv.ParseUint(idString, 10, 32)

	payload.ID = uint(id)

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
	payload.UpdatedBy = currentUserID

	res, err := handler.BackendSkeleton.Usecase.UpdateMenu(r.Context(), payload)
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, res, "Update Menu Success", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) UpdateOrderMenu(rw http.ResponseWriter, r *http.Request) {
	var payload request.UpdateOrderMenuRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		handler.WriteError(r.Context(), rw, lib.ErrorInvalidParameter)
		return
	}

	err = validation.Validator.Struct(payload)
	if err != nil {
		handler.WriteError(r.Context(), rw, lib.ErrorInvalidParameter)
		return
	}

	currentUserID := r.Context().Value("CurrentUserID").(uint)
	payload.UpdatedBy = currentUserID

	err = handler.BackendSkeleton.Usecase.UpdateOrderMenu(r.Context(), payload)
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, response.EmptyResponse{}, "Update Order Menu Success", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) DeleteMenu(rw http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "ID")
	id, _ := strconv.ParseUint(idString, 10, 32)

	err := handler.BackendSkeleton.Usecase.DeleteMenu(r.Context(), uint(id))
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, nil, "Delete Menu Success", ResponseMeta{HTTPStatus: http.StatusOK})
}
