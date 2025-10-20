package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Transaction(ctx context.Context, fn func(context.Context) error) error {
	arguments := m.Called(fn)
	if err := fn(ctx); err != nil {
		return err
	}
	return arguments.Error(0)
}

func (m *MockRepository) SendEmail(ctx context.Context, req lib.SMTPRequest) error {
	arguments := m.Called(req)
	return arguments.Error(0)
}
