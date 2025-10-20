package interface_repo

import (
	"context"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
)

type UserVerificationRepository interface {
	CreateUserVerification(ctx context.Context, user model.UserVerification) (model.UserVerification, error)
	GetUserVerificationByUserId(ctx context.Context, UserId uint) (model.UserVerification, error)
	UpdateVerificationStatusToTrue(ctx context.Context, ID uint) error
}
