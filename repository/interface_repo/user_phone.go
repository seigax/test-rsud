package interface_repo

import (
	"context"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

type UserPhoneRepository interface {
	CreateUserPhone(ctx context.Context, userPhone model.UserPhone) (model.UserPhone, error)
	GetUserPhones(ctx context.Context, userID uint) ([]model.UserPhone, error)
	UpdateUserPhone(ctx context.Context, userPhone model.UserPhone) (model.UserPhone, error)
	GetUserPhoneByID(ctx context.Context, ID uint) (model.UserPhone, error)
	GetUserPhone(ctx context.Context, phoneNumber string) (model.UserPhone, error)
	CheckPhoneNumberIsUsed(ctx context.Context, phoneNumber string) (bool, error)
	DeleteUserPhone(ctx context.Context, ID uint) error
	GetAllUserPhones(ctx context.Context, userID uint) ([]model.UserPhone, error)
	ChangeActiveUserPhone(ctx context.Context, req request.ReqUpdatePhone) error
}
