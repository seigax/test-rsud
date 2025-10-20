package interface_repo

import (
	"context"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
)

type BaseRepository interface {
	Transaction(ctx context.Context, fn func(context.Context) error) error
	SendEmail(ctx context.Context, req lib.SMTPRequest) error

	RedisRepository
	UserRepository
	ProvinceRepository
	CityRepository
	DistrictRepository
	VillageRepository
	RoleRepository
	UserRoleRepository
	UserPhoneRepository
	SessionRepository
	SystemParameterRepository
	ErrorMessageRepository
	MenuRepository
	UserVerificationRepository
	RoleMenuRepository
}
