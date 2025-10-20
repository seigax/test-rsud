package usecase

import (
	"context"
	"time"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/response"
)

func (usecase *Usecase) CreateErrorMessage(ctx context.Context, payload request.CreateErrorMessageRequest) (model.ErrorMessage, error) {
	ErrorMessage := model.ErrorMessage{
		Code:            payload.Code,
		Type:            payload.Type,
		ApplicationName: payload.ApplicationName,
		Message:         payload.Message,
		IsActiveFlag:    payload.IsActiveFlag,
		CreatedAt:       time.Now(),
		CreatedBy:       payload.CreatedBy,
		UpdatedAt:       time.Now(),
		UpdatedBy:       payload.CreatedBy,
	}

	res, err := usecase.repo.CreateErrorMessage(ctx, ErrorMessage)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (usecase *Usecase) GetErrorMessages(ctx context.Context, query request.GetErrorMessageQuery) (response.GetErrorMessageResponse, error) {
	var res response.GetErrorMessageResponse

	errorMessages, err := usecase.repo.GetErrorMessages(ctx, query)
	if err != nil {
		return res, err
	}

	total, err := usecase.repo.GetErrorMessageTotal(ctx, query)
	if err != nil {
		return res, err
	}

	res.ErrorMessages = errorMessages
	res.Total = total
	res.Limit = query.Limit
	res.Offset = query.GetOffset()
	res.Page = query.Page

	return res, nil
}

func (usecase *Usecase) GetErrorMessageDetail(ctx context.Context, ID uint) (model.ErrorMessage, error) {
	errorMessages, err := usecase.repo.GetErrorMessage(ctx, request.GetErrorMessageQuery{ErrorMessage: model.ErrorMessage{ID: ID}})
	if err != nil {
		return errorMessages, err
	}

	return errorMessages, nil
}

func (usecase *Usecase) UpdateErrorMessage(ctx context.Context, payload request.UpdateErrorMessageRequest) (model.ErrorMessage, error) {
	errorMessage, err := usecase.repo.GetErrorMessage(ctx, request.GetErrorMessageQuery{ErrorMessage: model.ErrorMessage{ID: payload.ID}})
	if err != nil {
		return errorMessage, err
	}

	errorMessage.Code = payload.Code
	errorMessage.Type = payload.Type
	errorMessage.ApplicationName = payload.ApplicationName
	errorMessage.Message = payload.Message
	errorMessage.IsActiveFlag = payload.IsActiveFlag
	errorMessage.UpdatedAt = time.Now()
	errorMessage.UpdatedBy = payload.UpdatedBy

	res, err := usecase.repo.UpdateErrorMessage(ctx, errorMessage)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (usecase *Usecase) DeleteErrorMessage(ctx context.Context, ID uint) error {
	err := usecase.repo.DeleteErrorMessage(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *Usecase) GetErrorMessageByCode(ctx context.Context, code string) (model.ErrorMessage, error) {
	errorMessages, err := usecase.repo.GetErrorMessage(ctx, request.GetErrorMessageQuery{ErrorMessage: model.ErrorMessage{Code: code}})
	if err != nil {
		return errorMessages, lib.ErrorInternalServer
	}

	return errorMessages, nil
}

func (usecase *Usecase) GetErrorMessageByCodeAndAppName(ctx context.Context, code string, appName string) (model.ErrorMessage, error) {
	errorMessages, err := usecase.repo.GetErrorMessage(ctx, request.GetErrorMessageQuery{ErrorMessage: model.ErrorMessage{Code: code, ApplicationName: appName}})
	if err != nil {
		return errorMessages, lib.ErrorInternalServer
	}

	return errorMessages, nil
}
