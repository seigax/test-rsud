package mock

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

func (m *MockRepository) CreateVillage(ctx context.Context, req model.Village) (model.Village, error) {
	arguments := m.Called(req)
	user := arguments.Get(0).(model.Village)

	if arguments.Error(1) != nil {
		return model.Village{}, arguments.Error(1)
	}

	return user, nil
}

func (m *MockRepository) GetVillages(ctx context.Context, query request.GetVillageQuery) ([]model.Village, error) {
	arguments := m.Called(query)

	if arguments.Error(1) != nil {
		return nil, arguments.Error(1)
	} else {
		villages := arguments.Get(0).([]model.Village)
		return villages, nil
	}
}

func (m *MockRepository) GetVillageTotal(ctx context.Context, query request.GetVillageQuery) (uint, error) {
	arguments := m.Called(query)

	if arguments.Error(1) != nil {
		return uint(0), arguments.Error(1)
	} else {
		total := arguments.Get(0).(uint)
		return total, nil
	}
}

func (m *MockRepository) UpdateVillage(ctx context.Context, village model.Village) (model.Village, error) {
	arguments := m.Called(village)

	if arguments.Error(1) != nil {
		return model.Village{}, arguments.Error(1)
	} else {
		village := arguments.Get(0).(model.Village)
		return village, nil
	}
}

func (m *MockRepository) GetVillageByID(ctx context.Context, ID uint) (model.Village, error) {
	arguments := m.Called(ID)

	if arguments.Error(1) != nil {
		return model.Village{}, arguments.Error(1)
	} else {
		village := arguments.Get(0).(model.Village)
		return village, nil
	}
}

func (m *MockRepository) GetVillagesByDistrictID(ctx context.Context, districtID uint) ([]model.Village, error) {
	arguments := m.Called(districtID)

	if arguments.Error(1) != nil {
		return []model.Village{}, arguments.Error(1)
	} else {
		village := arguments.Get(0).([]model.Village)
		return village, nil
	}
}

func (m *MockRepository) DeleteVillage(ctx context.Context, ID uint) error {
	arguments := m.Called(ID)
	err := arguments.Error(0)

	return err
}
