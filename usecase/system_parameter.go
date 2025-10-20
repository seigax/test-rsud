package usecase

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/response"
)

func (usecase *Usecase) CreateSystemParameter(ctx context.Context, payload request.CreateSystemParameterRequest) (systemParameter model.SystemParameter, err error) {
	err = usecase.repo.Transaction(ctx, func(ctx context.Context) error {
		systemParameter = model.SystemParameter{
			ParameterName: payload.ParameterName,
			DataType:      payload.DataType,
			Message:       payload.Message,
			IsActiveFlag:  payload.IsActiveFlag,
			CreatedAt:     time.Now(),
			CreatedBy:     payload.CreatedBy,
			UpdatedAt:     time.Now(),
			UpdatedBy:     payload.CreatedBy,
		}

		systemParameter, err = usecase.repo.CreateSystemParameter(ctx, systemParameter)
		if err != nil {
			return err
		}

		systemParameter.Code = fmt.Sprintf("SYS-%08d", systemParameter.ID)

		systemParameter, err = usecase.repo.UpdateSystemParameter(ctx, systemParameter)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

func (usecase *Usecase) GetSystemParameters(ctx context.Context, query request.GetSystemParameterQuery) (response.GetSystemParameterResponse, error) {
	var res response.GetSystemParameterResponse

	SystemParameters, err := usecase.repo.GetSystemParameters(ctx, query)
	if err != nil {
		return res, err
	}

	total, err := usecase.repo.GetSystemParameterTotal(ctx, query)
	if err != nil {
		return res, err
	}

	res.SystemParameters = SystemParameters
	res.Total = total
	res.Limit = query.Limit
	res.Offset = query.GetOffset()
	res.Page = query.Page

	return res, nil
}

func (usecase *Usecase) GetSystemParameterDetail(ctx context.Context, ID uint) (model.SystemParameter, error) {
	SystemParameters, err := usecase.repo.GetSystemParameter(ctx, request.GetSystemParameterQuery{SystemParameter: model.SystemParameter{ID: ID}})
	if err != nil {
		return SystemParameters, err
	}

	return SystemParameters, nil
}

func (usecase *Usecase) UpdateSystemParameter(ctx context.Context, payload request.UpdateSystemParameterRequest) (model.SystemParameter, error) {
	SystemParameter, err := usecase.repo.GetSystemParameter(ctx, request.GetSystemParameterQuery{SystemParameter: model.SystemParameter{ID: payload.ID}})
	if err != nil {
		return SystemParameter, err
	}

	SystemParameter.ParameterName = payload.ParameterName
	SystemParameter.DataType = payload.DataType
	SystemParameter.Message = payload.Message
	SystemParameter.IsActiveFlag = payload.IsActiveFlag
	SystemParameter.UpdatedAt = time.Now()
	SystemParameter.UpdatedBy = payload.UpdatedBy

	res, err := usecase.repo.UpdateSystemParameter(ctx, SystemParameter)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (usecase *Usecase) DeleteSystemParameter(ctx context.Context, ID uint) error {
	err := usecase.repo.DeleteSystemParameter(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
