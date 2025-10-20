package mock

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

func (m *MockRepository) CreateCity(ctx context.Context, req model.City) (model.City, error) {
	arguments := m.Called(req)
	user := arguments.Get(0).(model.City)

	if arguments.Error(1) != nil {
		return model.City{}, arguments.Error(1)
	}

	return user, nil
}

func (m *MockRepository) GetCitys(ctx context.Context, query request.GetCityQuery) ([]model.City, error) {
	arguments := m.Called(query)
	if arguments.Error(1) != nil {
		return nil, arguments.Error(1)
	} else {
		Citys := arguments.Get(0).([]model.City)
		return Citys, nil
	}
}

func (m *MockRepository) GetCityTotal(ctx context.Context, query request.GetCityQuery) (uint, error) {
	arguments := m.Called(query)

	if arguments.Error(1) != nil {
		return uint(0), arguments.Error(1)
	} else {
		total := arguments.Get(0).(uint)
		return total, nil
	}
}

func (m *MockRepository) UpdateCity(ctx context.Context, city model.City) (model.City, error) {
	arguments := m.Called(city)

	if arguments.Error(1) != nil {
		return model.City{}, arguments.Error(1)
	} else {
		City := arguments.Get(0).(model.City)
		return City, nil
	}
}

func (m *MockRepository) GetCity(ctx context.Context, query request.GetCityQuery) (model.City, error) {
	arguments := m.Called(query)

	if arguments.Error(1) != nil {
		return model.City{}, arguments.Error(1)
	} else {
		City := arguments.Get(0).(model.City)
		return City, nil
	}
}

func (m *MockRepository) DeleteCity(ctx context.Context, ID uint) error {
	arguments := m.Called(ID)
	err := arguments.Error(0)

	return err
}
