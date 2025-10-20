package interface_repo

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	GetUsers(ctx context.Context, query request.GetUserQuery) ([]model.User, error)
	GetUserTotal(ctx context.Context, query request.GetUserQuery) (uint, error)
	UpdateUser(ctx context.Context, user model.User) (model.User, error)
	GetUserByID(ctx context.Context, ID uint) (model.User, error)
	GetUserByEmail(ctx context.Context, email string) (user model.User, err error)
	DeleteUser(ctx context.Context, ID uint) error
}
