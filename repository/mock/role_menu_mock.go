package mock

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
)

func (m *MockRepository) CreateRoleMenu(ctx context.Context, req model.RoleMenu) (model.RoleMenu, error) {
	arguments := m.Called(req)
	user := arguments.Get(0).(model.RoleMenu)

	if arguments.Error(1) != nil {
		return model.RoleMenu{}, arguments.Error(1)
	}

	return user, nil
}

func (m *MockRepository) GetRoleMenusByRoleID(ctx context.Context, roleID uint) ([]model.RoleMenu, error) {
	arguments := m.Called(roleID)

	if arguments.Error(1) != nil {
		return []model.RoleMenu{}, arguments.Error(1)
	} else {
		roleMenu := arguments.Get(0).([]model.RoleMenu)
		return roleMenu, nil
	}
}

func (m *MockRepository) GetRoleMenusWithRoleUrlByRoleID(ctx context.Context, roleID uint) ([]model.RoleMenuWithRoleUrl, error) {
	arguments := m.Called(roleID)

	if arguments.Error(1) != nil {
		return []model.RoleMenuWithRoleUrl{}, arguments.Error(1)
	} else {
		roleMenu := arguments.Get(0).([]model.RoleMenuWithRoleUrl)
		return roleMenu, nil
	}
}

func (m *MockRepository) GetRoleMenuByRoleIDAndMenuID(ctx context.Context, roleID uint, menuID uint) (model.RoleMenu, error) {
	arguments := m.Called(roleID)

	if arguments.Error(1) != nil {
		return model.RoleMenu{}, arguments.Error(1)
	} else {
		roleMenu := arguments.Get(0).(model.RoleMenu)
		return roleMenu, nil
	}
}

func (m *MockRepository) DeleteRoleMenu(ctx context.Context, ID uint) error {
	arguments := m.Called(ID)
	err := arguments.Error(0)

	return err
}

func (m *MockRepository) DeleteAllRoleMenu(ctx context.Context, roleID uint) error {
	arguments := m.Called(roleID)
	err := arguments.Error(0)

	return err
}
