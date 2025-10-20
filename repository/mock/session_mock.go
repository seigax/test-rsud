package mock

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
)

func (m *MockRepository) CreateSession(ctx context.Context, req model.Session) (model.Session, error) {
	arguments := m.Called(req)
	user := arguments.Get(0).(model.Session)

	if arguments.Error(1) != nil {
		return model.Session{}, arguments.Error(1)
	}

	return user, nil
}

func (m *MockRepository) UpdateSession(ctx context.Context, session model.Session) (model.Session, error) {
	arguments := m.Called(session)

	if arguments.Error(1) != nil {
		return model.Session{}, arguments.Error(1)
	} else {
		session := arguments.Get(0).(model.Session)
		return session, nil
	}
}

func (m *MockRepository) GetSession(ctx context.Context, session model.Session) (model.Session, error) {
	arguments := m.Called(session)

	if arguments.Error(1) != nil {
		return model.Session{}, arguments.Error(1)
	} else {
		session := arguments.Get(0).(model.Session)
		return session, nil
	}
}

func (m *MockRepository) DeleteSession(ctx context.Context, ID uint) error {
	arguments := m.Called(ID)
	err := arguments.Error(0)

	return err
}
