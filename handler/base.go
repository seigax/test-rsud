package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	backendskeleton "gitlab.com/erloom.id/libraries/go/backend-skeleton"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/validation"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/response"
)

type Handler struct {
	BackendSkeleton *backendskeleton.BackendSkeleton
}

func NewHandler(backendskeleton *backendskeleton.BackendSkeleton) Handler {
	return Handler{
		BackendSkeleton: backendskeleton,
	}
}

func (handler *Handler) Healthz(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(rw, "server is ok")
}

type ResponseBody struct {
	Data    interface{}  `json:"data,omitempty"`
	Message string       `json:"message,omitempty"`
	Meta    ResponseMeta `json:"meta"`
}

type ResponseMeta struct {
	HTTPStatus   int         `json:"http_status"`
	Total        *uint       `json:"total,omitempty"`
	Offset       *uint       `json:"offset,omitempty"`
	Limit        *uint       `json:"limit,omitempty"`
	Page         *uint       `json:"page,omitempty"`
	LastPage     *uint       `json:"last_page,omitempty"`
	ValidatorErr interface{} `json:"error_validation,omitempty"`
}

type ErrorInfo struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}

type ErrorBody struct {
	Errors []ErrorInfo `json:"errors"`
	Meta   interface{} `json:"meta"`
}

func WriteError(w http.ResponseWriter, err error) {
	var resp interface{}
	code := http.StatusInternalServerError

	switch errOrig := err.(type) {
	case lib.CustomError:
		resp = ErrorBody{
			Errors: []ErrorInfo{
				{
					Message: errOrig.Message,
					Field:   errOrig.Field,
				},
			},
			Meta: ResponseMeta{
				HTTPStatus: errOrig.HTTPCode,
			},
		}

		code = errOrig.HTTPCode
	default:
		resp = ResponseBody{
			Message: "Internal Server Error",
			Meta: ResponseMeta{
				HTTPStatus: code,
			},
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

func (handler *Handler) WriteError(ctx context.Context, w http.ResponseWriter, err error) {
	var resp interface{}
	httpCode := http.StatusInternalServerError

	requestAppName := ctx.Value("X-Request-AppName").(string)

	switch errOrig := err.(type) {
	case lib.CustomError:
		errMsg, err := handler.BackendSkeleton.Usecase.GetErrorMessageByCodeAndAppName(ctx, fmt.Sprint(errOrig.Code), requestAppName)
		if err != nil {
			resp = ResponseBody{
				Message: "Internal Server Error",
				Meta: ResponseMeta{
					HTTPStatus: httpCode,
				},
			}
			break
		}

		resp = ErrorBody{
			Errors: []ErrorInfo{
				{
					Message: errMsg.Message,
					Field:   errOrig.Field,
				},
			},
			Meta: ResponseMeta{
				HTTPStatus: errOrig.HTTPCode,
			},
		}

		httpCode = errOrig.HTTPCode
	case validator.ValidationErrors:
		errMsg, err := handler.BackendSkeleton.Usecase.GetErrorMessageByCodeAndAppName(ctx, fmt.Sprint(lib.ErrorInvalidParameter.Code), requestAppName)
		httpCode = lib.ErrorInvalidParameter.HTTPCode
		if err != nil {
			resp = ResponseBody{
				Message: lib.ErrorInvalidParameter.Message,
				Meta: ResponseMeta{
					HTTPStatus: httpCode,
				},
			}
			break
		}

		errInfo := []ErrorInfo{}
		for _, v := range errOrig {
			lang := ctx.Value("RequestLang").(string)
			translatedMsg := ""
			switch lang {
			case "en":
				translatedMsg = v.Translate(validation.ENTrans)
			default:
				translatedMsg = v.Translate(validation.IDTrans)
			}
			errInfo = append(errInfo, ErrorInfo{
				Message: translatedMsg,
				Field:   v.Field(),
			})
		}
		resp = ErrorBody{
			Errors: []ErrorInfo{
				{
					Message: errMsg.Message,
				},
			},
			Meta: ResponseMeta{
				HTTPStatus:   http.StatusUnprocessableEntity,
				ValidatorErr: errInfo,
			},
		}
	default:
		resp = ResponseBody{
			Message: "Internal Server Error",
			Meta: ResponseMeta{
				HTTPStatus: httpCode,
			},
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(resp)
}

func WriteSuccess(w http.ResponseWriter, data interface{}, message string, meta ResponseMeta) {
	resp := ResponseBody{
		Message: message,
		Data:    data,
		Meta:    meta,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(meta.HTTPStatus)
	json.NewEncoder(w).Encode(resp)
}

func WriteResponse(w http.ResponseWriter, resp interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func (handler *ResponseMeta) SerializeFromResponse(resp response.BasePaginateResponse) {
	if resp.Limit == 0 {
		return
	}
	handler.Total = &resp.Total
	handler.Page = &resp.Page

	lastPage := resp.Total / resp.Limit
	if (lastPage * resp.Limit) < resp.Total {
		lastPage++
	}
	if lastPage == 0 {
		lastPage++
	}

	handler.LastPage = &lastPage
}
