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

func (handler *Handler) CreateVillage(rw http.ResponseWriter, r *http.Request) {
	var payload request.CreateVillageRequest

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

	res, err := handler.BackendSkeleton.Usecase.CreateVillage(r.Context(), payload)
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, res, "Create Village Success", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) GetVillages(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	var limit, page, districtId int

	limit, _ = strconv.Atoi(queryParams.Get("limit"))
	page, _ = strconv.Atoi(queryParams.Get("page"))
	districtId, _ = strconv.Atoi(queryParams.Get("district_id"))

	sort := strings.Split(queryParams.Get("sort"), ",")

	var query request.GetVillageQuery

	query.Limit = uint(limit)
	query.Page = uint(page)
	query.Search = queryParams.Get("search")
	query.Sort = sort
	query.DistrictID = uint(districtId)

	res, err := handler.BackendSkeleton.Usecase.GetVillages(r.Context(), query)
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	resp := ResponseMeta{HTTPStatus: http.StatusOK}
	resp.SerializeFromResponse(res.BasePaginateResponse)

	WriteSuccess(rw, res, "success", resp)
}

func (handler *Handler) GetVillageDetail(rw http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "ID")
	id, _ := strconv.ParseUint(idString, 10, 32)

	res, err := handler.BackendSkeleton.Usecase.GetVillageDetail(r.Context(), uint(id))
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, res, "success", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) UpdateVillage(rw http.ResponseWriter, r *http.Request) {
	var payload request.UpdateVillageRequest

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

	res, err := handler.BackendSkeleton.Usecase.UpdateVillage(r.Context(), payload)
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, res, "Update Village Success", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) DeleteVillage(rw http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "ID")
	id, _ := strconv.ParseUint(idString, 10, 32)

	err := handler.BackendSkeleton.Usecase.DeleteVillage(r.Context(), uint(id))
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, nil, "Delete Village Success", ResponseMeta{HTTPStatus: http.StatusOK})
}
