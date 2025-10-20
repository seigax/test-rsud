package mock

import (
	"context"
	"time"
)

func (m *MockRepository) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	arguments := m.Called(key, value, expiration)

	if arguments.Error(1) != nil {
		return arguments.Error(1)
	}

	return nil
}

func (m *MockRepository) Get(ctx context.Context, key string) (string, error) {
	arguments := m.Called(key)

	if arguments.Error(1) != nil {
		return "", arguments.Error(1)
	}

	data := arguments.Get(0).(string)
	return data, nil
}

func (m *MockRepository) Delete(ctx context.Context, key string) error {
	arguments := m.Called(key)

	if arguments.Error(1) != nil {
		return arguments.Error(1)
	}

	return nil
}
