package mock

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

func (m *MockRepository) CreateMenu(ctx context.Context, req model.Menu) (model.Menu, error) {
	arguments := m.Called(req)
	user := arguments.Get(0).(model.Menu)

	if arguments.Error(1) != nil {
		return model.Menu{}, arguments.Error(1)
	}

	return user, nil
}

func (m *MockRepository) GetMenus(ctx context.Context, query request.GetMenuQuery) ([]model.Menu, error) {
	arguments := m.Called(query)
	if arguments.Error(1) != nil {
		return nil, arguments.Error(1)
	} else {
		menus := arguments.Get(0).([]model.Menu)
		return menus, nil
	}
}

func (m *MockRepository) GetMenuTotal(ctx context.Context, query request.GetMenuQuery) (uint, error) {
	arguments := m.Called(query)

	if arguments.Error(1) != nil {
		return uint(0), arguments.Error(1)
	} else {
		total := arguments.Get(0).(uint)
		return total, nil
	}
}

func (m *MockRepository) UpdateMenu(ctx context.Context, menu model.Menu) (model.Menu, error) {
	arguments := m.Called(menu)

	if arguments.Error(1) != nil {
		return model.Menu{}, arguments.Error(1)
	} else {
		menu := arguments.Get(0).(model.Menu)
		return menu, nil
	}
}

func (m *MockRepository) GetMenu(ctx context.Context, query request.GetMenuQuery) (model.Menu, error) {
	arguments := m.Called()

	if arguments.Error(1) != nil {
		return model.Menu{}, arguments.Error(1)
	} else {
		menu := arguments.Get(0).(model.Menu)
		return menu, nil
	}
}

func (m *MockRepository) DeleteMenu(ctx context.Context, ID uint) error {
	arguments := m.Called(ID)
	err := arguments.Error(0)

	return err
}
