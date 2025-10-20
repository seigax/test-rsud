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

func (handler *Handler) CreateDistrict(rw http.ResponseWriter, r *http.Request) {
	var payload request.CreateDistrictRequest

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

	res, err := handler.BackendSkeleton.Usecase.CreateDistrict(r.Context(), payload)
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, res, "Create District Success", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) GetDistricts(rw http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	var limit, page, city_id int

	limit, _ = strconv.Atoi(queryParams.Get("limit"))
	page, _ = strconv.Atoi(queryParams.Get("page"))
	city_id, _ = strconv.Atoi(queryParams.Get("city_id"))

	sort := strings.Split(queryParams.Get("sort"), ",")

	var query request.GetDistrictQuery

	query.Limit = uint(limit)
	query.Page = uint(page)
	query.Search = queryParams.Get("search")
	query.Sort = sort
	query.CityID = uint(city_id)

	res, err := handler.BackendSkeleton.Usecase.GetDistricts(r.Context(), query)
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	resp := ResponseMeta{HTTPStatus: http.StatusOK}
	resp.SerializeFromResponse(res.BasePaginateResponse)

	WriteSuccess(rw, res, "success", resp)
}

func (handler *Handler) GetDistrictDetail(rw http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "ID")
	id, _ := strconv.ParseUint(idString, 10, 32)

	res, err := handler.BackendSkeleton.Usecase.GetDistrictDetail(r.Context(), uint(id))
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, res, "success", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) UpdateDistrict(rw http.ResponseWriter, r *http.Request) {
	var payload request.UpdateDistrictRequest

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

	res, err := handler.BackendSkeleton.Usecase.UpdateDistrict(r.Context(), payload)
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, res, "Update District Success", ResponseMeta{HTTPStatus: http.StatusOK})
}

func (handler *Handler) DeleteDistrict(rw http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "ID")
	id, _ := strconv.ParseUint(idString, 10, 32)

	err := handler.BackendSkeleton.Usecase.DeleteDistrict(r.Context(), uint(id))
	if err != nil {
		handler.WriteError(r.Context(), rw, err)
		return
	}

	WriteSuccess(rw, nil, "Delete District Success", ResponseMeta{HTTPStatus: http.StatusOK})
}
