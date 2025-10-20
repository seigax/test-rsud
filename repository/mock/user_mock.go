package mock

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

func (m *MockRepository) CreateUser(ctx context.Context, req model.User) (model.User, error) {
	arguments := m.Called(req)
	user := arguments.Get(0).(model.User)

	if arguments.Error(1) != nil {
		return model.User{}, arguments.Error(1)
	}

	return user, nil
}

func (m *MockRepository) GetUsers(ctx context.Context, query request.GetUserQuery) ([]model.User, error) {
	arguments := m.Called(query)

	if arguments.Error(1) != nil {
		return nil, arguments.Error(1)
	} else {
		users := arguments.Get(0).([]model.User)
		return users, nil
	}
}

func (m *MockRepository) GetUserTotal(ctx context.Context, query request.GetUserQuery) (uint, error) {
	arguments := m.Called(query)

	if arguments.Error(1) != nil {
		return uint(0), arguments.Error(1)
	} else {
		total := arguments.Get(0).(uint)
		return total, nil
	}
}

func (m *MockRepository) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	arguments := m.Called(user)

	if arguments.Error(1) != nil {
		return model.User{}, arguments.Error(1)
	} else {
		user := arguments.Get(0).(model.User)
		return user, nil
	}
}

func (m *MockRepository) GetUserByID(ctx context.Context, ID uint) (model.User, error) {
	arguments := m.Called(ID)

	if arguments.Error(1) != nil {
		return model.User{}, arguments.Error(1)
	} else {
		user := arguments.Get(0).(model.User)
		return user, nil
	}
}

func (m *MockRepository) GetUserByEmail(ctx context.Context, email string) (user model.User, err error) {
	arguments := m.Called(email)
	if arguments.Error(1) != nil {
		return model.User{}, arguments.Error(1)
	} else {
		user := arguments.Get(0).(model.User)
		return user, nil
	}
}

func (m *MockRepository) DeleteUser(ctx context.Context, ID uint) error {
	arguments := m.Called(ID)
	err := arguments.Error(0)

	return err
}
