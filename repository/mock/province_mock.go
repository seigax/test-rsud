package mock

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

func (m *MockRepository) CreateProvince(ctx context.Context, req model.Province) (model.Province, error) {
	arguments := m.Called(req)
	user := arguments.Get(0).(model.Province)

	if arguments.Error(1) != nil {
		return model.Province{}, arguments.Error(1)
	}

	return user, nil
}

func (m *MockRepository) GetProvinces(ctx context.Context, query request.GetProvinceQuery) ([]model.Province, error) {
	arguments := m.Called(query)

	if arguments.Error(1) != nil {
		return nil, arguments.Error(1)
	} else {
		provinces := arguments.Get(0).([]model.Province)
		return provinces, nil
	}
}

func (m *MockRepository) GetProvinceTotal(ctx context.Context, query request.GetProvinceQuery) (uint, error) {
	arguments := m.Called(query)

	if arguments.Error(1) != nil {
		return uint(0), arguments.Error(1)
	} else {
		total := arguments.Get(0).(uint)
		return total, nil
	}
}

func (m *MockRepository) UpdateProvince(ctx context.Context, province model.Province) (model.Province, error) {
	arguments := m.Called(province)

	if arguments.Error(1) != nil {
		return model.Province{}, arguments.Error(1)
	} else {
		province := arguments.Get(0).(model.Province)
		return province, nil
	}
}

func (m *MockRepository) UpdateProvinceAddTotalCity(ctx context.Context, provinceID uint, add int) error {
	arguments := m.Called(provinceID)

	if arguments.Error(0) != nil {
		return arguments.Error(0)
	} else {
		return nil
	}
}

func (m *MockRepository) UpdateProvinceAddTotalDistrict(ctx context.Context, provinceID uint, add int) error {
	arguments := m.Called(provinceID)

	if arguments.Error(0) != nil {
		return arguments.Error(0)
	} else {
		return nil
	}
}

func (m *MockRepository) UpdateProvinceAddTotalVillage(ctx context.Context, provinceID uint, add int) error {
	arguments := m.Called(provinceID)

	if arguments.Error(0) != nil {
		return arguments.Error(0)
	} else {
		return nil
	}
}

func (m *MockRepository) GetProvince(ctx context.Context, query request.GetProvinceQuery) (province model.Province, err error) {
	arguments := m.Called(query)

	if arguments.Error(1) != nil {
		return model.Province{}, arguments.Error(1)
	} else {
		province := arguments.Get(0).(model.Province)
		return province, nil
	}
}

func (m *MockRepository) DeleteProvince(ctx context.Context, ID uint) error {
	arguments := m.Called(ID)
	err := arguments.Error(0)

	return err
}
