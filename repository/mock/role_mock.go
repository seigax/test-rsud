package mock

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

func (m *MockRepository) CreateRole(ctx context.Context, req model.Role) (model.Role, error) {
	arguments := m.Called(req)
	user := arguments.Get(0).(model.Role)

	if arguments.Error(1) != nil {
		return model.Role{}, arguments.Error(1)
	}

	return user, nil
}

func (m *MockRepository) GetRoles(ctx context.Context, query request.GetRoleQuery) ([]model.Role, error) {
	arguments := m.Called(query)

	if arguments.Error(1) != nil {
		return nil, arguments.Error(1)
	} else {
		roles := arguments.Get(0).([]model.Role)
		return roles, nil
	}
}

func (m *MockRepository) GetRoleTotal(ctx context.Context, query request.GetRoleQuery) (uint, error) {
	arguments := m.Called(query)

	if arguments.Error(1) != nil {
		return uint(0), arguments.Error(1)
	} else {
		total := arguments.Get(0).(uint)
		return total, nil
	}
}

func (m *MockRepository) UpdateRole(ctx context.Context, role model.Role) (model.Role, error) {
	arguments := m.Called(role)

	if arguments.Error(1) != nil {
		return model.Role{}, arguments.Error(1)
	} else {
		role := arguments.Get(0).(model.Role)
		return role, nil
	}
}

func (m *MockRepository) GetRole(ctx context.Context, query request.GetRoleQuery) (model.Role, error) {
	arguments := m.Called()

	if arguments.Error(1) != nil {
		return model.Role{}, arguments.Error(1)
	} else {
		role := arguments.Get(0).(model.Role)
		return role, nil
	}
}

func (m *MockRepository) DeleteRole(ctx context.Context, ID uint) error {
	arguments := m.Called(ID)
	err := arguments.Error(0)

	return err
}
