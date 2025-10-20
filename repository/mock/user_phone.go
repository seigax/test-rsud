package mock

import (
	"context"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
)

func (m *MockRepository) CreateUserPhone(ctx context.Context, req model.UserPhone) (model.UserPhone, error) {
	arguments := m.Called(req)
	user := arguments.Get(0).(model.UserPhone)

	if arguments.Error(1) != nil {
		return model.UserPhone{}, arguments.Error(1)
	}

	return user, nil
}

func (m *MockRepository) GetUserPhones(ctx context.Context, userID uint) ([]model.UserPhone, error) {
	arguments := m.Called(userID)

	if arguments.Error(1) != nil {
		return nil, arguments.Error(1)
	} else {
		userPhone := arguments.Get(0).([]model.UserPhone)
		return userPhone, nil
	}
}

func (m *MockRepository) UpdateUserPhone(ctx context.Context, UserPhone model.UserPhone) (model.UserPhone, error) {
	arguments := m.Called(UserPhone)

	if arguments.Error(1) != nil {
		return model.UserPhone{}, arguments.Error(1)
	} else {
		userPhone := arguments.Get(0).(model.UserPhone)
		return userPhone, nil
	}
}

func (m *MockRepository) GetUserPhoneByID(ctx context.Context, ID uint) (model.UserPhone, error) {
	arguments := m.Called(ID)

	if arguments.Error(1) != nil {
		return model.UserPhone{}, arguments.Error(1)
	} else {
		userPhone := arguments.Get(0).(model.UserPhone)
		return userPhone, nil
	}
}

func (m *MockRepository) GetUserPhone(ctx context.Context, phoneNumber string) (model.UserPhone, error) {
	arguments := m.Called(phoneNumber)

	if arguments.Error(1) != nil {
		return model.UserPhone{}, arguments.Error(1)
	} else {
		userPhone := arguments.Get(0).(model.UserPhone)
		return userPhone, nil
	}
}

func (m *MockRepository) CheckPhoneNumberIsUsed(ctx context.Context, phoneNumber string) (bool, error) {
	arguments := m.Called(phoneNumber)
	return arguments.Bool(0), arguments.Error(1)
}

func (m *MockRepository) DeleteUserPhone(ctx context.Context, ID uint) error {
	arguments := m.Called(ID)
	err := arguments.Error(0)

	return err
}

func (m *MockRepository) GetAllUserPhones(ctx context.Context, userID uint) ([]model.UserPhone, error) {
	arguments := m.Called(userID)

	if arguments.Error(1) != nil {
		return nil, arguments.Error(1)
	} else {
		userPhone := arguments.Get(0).([]model.UserPhone)
		return userPhone, nil
	}
}

func (m *MockRepository) ChangeActiveUserPhone(ctx context.Context, req request.ReqUpdatePhone) error {
	arguments := m.Called(req)
	err := arguments.Error(0)

	return err
}
