package mock

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

func (m *MockRepository) CreateDistrict(ctx context.Context, req model.District) (model.District, error) {
	arguments := m.Called(req)
	user := arguments.Get(0).(model.District)

	if arguments.Error(1) != nil {
		return model.District{}, arguments.Error(1)
	}

	return user, nil
}

func (m *MockRepository) GetDistricts(ctx context.Context, query request.GetDistrictQuery) ([]model.District, error) {
	arguments := m.Called(query)

	if arguments.Error(1) != nil {
		return nil, arguments.Error(1)
	} else {
		Districts := arguments.Get(0).([]model.District)
		return Districts, nil
	}
}

func (m *MockRepository) GetDistrictTotal(ctx context.Context, query request.GetDistrictQuery) (uint, error) {
	arguments := m.Called(query)

	if arguments.Error(1) != nil {
		return uint(0), arguments.Error(1)
	} else {
		total := arguments.Get(0).(uint)
		return total, nil
	}
}

func (m *MockRepository) UpdateDistrict(ctx context.Context, district model.District) (model.District, error) {
	arguments := m.Called(district)

	if arguments.Error(1) != nil {
		return model.District{}, arguments.Error(1)
	} else {
		district := arguments.Get(0).(model.District)
		return district, nil
	}
}

func (m *MockRepository) GetDistrict(ctx context.Context, query request.GetDistrictQuery) (model.District, error) {
	arguments := m.Called()

	if arguments.Error(1) != nil {
		return model.District{}, arguments.Error(1)
	} else {
		District := arguments.Get(0).(model.District)
		return District, nil
	}
}

func (m *MockRepository) DeleteDistrict(ctx context.Context, ID uint) error {
	arguments := m.Called(ID)
	err := arguments.Error(0)

	return err
}
