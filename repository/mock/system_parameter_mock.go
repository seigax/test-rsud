package mock

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

func (m *MockRepository) CreateSystemParameter(ctx context.Context, req model.SystemParameter) (model.SystemParameter, error) {
	arguments := m.Called(req)
	user := arguments.Get(0).(model.SystemParameter)

	if arguments.Error(1) != nil {
		return model.SystemParameter{}, arguments.Error(1)
	}

	return user, nil
}

func (m *MockRepository) GetSystemParameters(ctx context.Context, query request.GetSystemParameterQuery) ([]model.SystemParameter, error) {
	arguments := m.Called(query)
	if arguments.Error(1) != nil {
		return nil, arguments.Error(1)
	} else {
		SystemParameters := arguments.Get(0).([]model.SystemParameter)
		return SystemParameters, nil
	}
}

func (m *MockRepository) GetSystemParameterTotal(ctx context.Context, query request.GetSystemParameterQuery) (uint, error) {
	arguments := m.Called(query)

	if arguments.Error(1) != nil {
		return uint(0), arguments.Error(1)
	} else {
		total := arguments.Get(0).(uint)
		return total, nil
	}
}

func (m *MockRepository) UpdateSystemParameter(ctx context.Context, SystemParameters model.SystemParameter) (model.SystemParameter, error) {
	arguments := m.Called(SystemParameters)

	if arguments.Error(1) != nil {
		return model.SystemParameter{}, arguments.Error(1)
	} else {
		SystemParameters := arguments.Get(0).(model.SystemParameter)
		return SystemParameters, nil
	}
}

func (m *MockRepository) GetSystemParameter(ctx context.Context, query request.GetSystemParameterQuery) (model.SystemParameter, error) {
	arguments := m.Called(query)

	if arguments.Error(1) != nil {
		return model.SystemParameter{}, arguments.Error(1)
	} else {
		SystemParameters := arguments.Get(0).(model.SystemParameter)
		return SystemParameters, nil
	}
}

func (m *MockRepository) DeleteSystemParameter(ctx context.Context, ID uint) error {
	arguments := m.Called(ID)
	err := arguments.Error(0)

	return err
}
