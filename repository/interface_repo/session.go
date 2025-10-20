package interface_repo

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, session model.Session) (model.Session, error)
	UpdateSession(ctx context.Context, session model.Session) (model.Session, error)
	GetSession(ctx context.Context, session model.Session) (model.Session, error)
	DeleteSession(ctx context.Context, ID uint) error
}
