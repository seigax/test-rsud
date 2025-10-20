package mock

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

func (m *MockRepository) CreateErrorMessage(ctx context.Context, req model.ErrorMessage) (model.ErrorMessage, error) {
	arguments := m.Called(req)
	user := arguments.Get(0).(model.ErrorMessage)

	if arguments.Error(1) != nil {
		return model.ErrorMessage{}, arguments.Error(1)
	}

	return user, nil
}

func (m *MockRepository) GetErrorMessages(ctx context.Context, query request.GetErrorMessageQuery) ([]model.ErrorMessage, error) {
	arguments := m.Called(query)
	if arguments.Error(1) != nil {
		return nil, arguments.Error(1)
	} else {
		errorMessages := arguments.Get(0).([]model.ErrorMessage)
		return errorMessages, nil
	}
}

func (m *MockRepository) GetErrorMessageTotal(ctx context.Context, query request.GetErrorMessageQuery) (uint, error) {
	arguments := m.Called(query)

	if arguments.Error(1) != nil {
		return uint(0), arguments.Error(1)
	} else {
		total := arguments.Get(0).(uint)
		return total, nil
	}
}

func (m *MockRepository) UpdateErrorMessage(ctx context.Context, errorMessages model.ErrorMessage) (model.ErrorMessage, error) {
	arguments := m.Called(errorMessages)

	if arguments.Error(1) != nil {
		return model.ErrorMessage{}, arguments.Error(1)
	} else {
		errorMessages := arguments.Get(0).(model.ErrorMessage)
		return errorMessages, nil
	}
}

func (m *MockRepository) GetErrorMessage(ctx context.Context, query request.GetErrorMessageQuery) (model.ErrorMessage, error) {
	arguments := m.Called()

	if arguments.Error(1) != nil {
		return model.ErrorMessage{}, arguments.Error(1)
	} else {
		errorMessages := arguments.Get(0).(model.ErrorMessage)
		return errorMessages, nil
	}
}

func (m *MockRepository) DeleteErrorMessage(ctx context.Context, ID uint) error {
	arguments := m.Called(ID)
	err := arguments.Error(0)

	return err
}
