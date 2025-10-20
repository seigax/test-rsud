package mock

import (
	"context"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
)

func (m *MockRepository) CreateUserVerification(ctx context.Context, user model.UserVerification) (model.UserVerification, error) {

	arguments := m.Called(user)
	result := arguments.Get(0).(model.UserVerification)

	if arguments.Error(1) != nil {
		return model.UserVerification{}, arguments.Error(1)
	}

	return result, nil

}

func (m *MockRepository) GetUserVerificationByUserId(ctx context.Context, UserId uint) (model.UserVerification, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockRepository) UpdateVerificationStatusToTrue(ctx context.Context, ID uint) error {
	//TODO implement me
	panic("implement me")
}
