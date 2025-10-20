package mock

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
)

func (m *MockRepository) CreateUserRole(ctx context.Context, req model.UserRole) (model.UserRole, error) {
	arguments := m.Called(req)
	user := arguments.Get(0).(model.UserRole)

	if arguments.Error(1) != nil {
		return model.UserRole{}, arguments.Error(1)
	}

	return user, nil
}

func (m *MockRepository) UpdateUserRole(ctx context.Context, userrole model.UserRole) (model.UserRole, error) {
	arguments := m.Called(userrole)

	if arguments.Error(1) != nil {
		return model.UserRole{}, arguments.Error(1)
	} else {
		userrole := arguments.Get(0).(model.UserRole)
		return userrole, nil
	}
}

func (m *MockRepository) GetUserRoleByID(ctx context.Context, ID uint) (model.UserRole, error) {
	arguments := m.Called(ID)

	if arguments.Error(1) != nil {
		return model.UserRole{}, arguments.Error(1)
	} else {
		userRole := arguments.Get(0).(model.UserRole)
		return userRole, nil
	}
}

func (m *MockRepository) GetUserRoleByUserID(ctx context.Context, userID uint) (model.UserRole, error) {
	arguments := m.Called(userID)

	if arguments.Error(1) != nil {
		return model.UserRole{}, arguments.Error(1)
	} else {
		userRole := arguments.Get(0).(model.UserRole)
		return userRole, nil
	}
}

func (m *MockRepository) GetUserRolesWithRoleDetailByUserID(ctx context.Context, userID uint) ([]model.UserRoleWithRoleDetail, error) {
	arguments := m.Called(userID)

	if arguments.Error(1) != nil {
		return []model.UserRoleWithRoleDetail{}, arguments.Error(1)
	} else {
		userRole := arguments.Get(0).([]model.UserRoleWithRoleDetail)
		return userRole, nil
	}
}

func (m *MockRepository) GetUserRolesByUserID(ctx context.Context, userID uint) ([]model.UserRole, error) {
	arguments := m.Called(userID)

	if arguments.Error(1) != nil {
		return []model.UserRole{}, arguments.Error(1)
	} else {
		userRole := arguments.Get(0).([]model.UserRole)
		return userRole, nil
	}
}

func (m *MockRepository) DeleteUserRole(ctx context.Context, ID uint) error {
	arguments := m.Called(ID)
	err := arguments.Error(0)

	return err
}

func (m *MockRepository) DeleteUserRoleByUserID(ctx context.Context, userID uint) error {
	arguments := m.Called(userID)
	err := arguments.Error(0)

	return err
}
