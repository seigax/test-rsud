package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/response"
)

// func (usecase *Usecase) RegisterFarmer(ctx context.Context, payload request.RegisterFarmerRequest) (model.User, error) {
// 	var user model.User

// 	checkEmail, err := usecase.repo.GetUserByEmail(ctx, payload.Email)
// 	if checkEmail.ID != 0 {
// 		return user, lib.ErrorEmailAlreadyRegistered
// 	}
// 	if err != nil && !errors.Is(lib.ErrorNotFound, err) {
// 		return user, lib.ErrorInternalServer
// 	}

// 	isPhoneUsed, err := usecase.repo.CheckPhoneNumberIsUsed(ctx, payload.Phone)
// 	if isPhoneUsed {
// 		return user, lib.ErrorPhoneAlreadyRegistered
// 	}
// 	if err != nil && !errors.Is(lib.ErrorNotFound, err) {
// 		return user, lib.ErrorInternalServer
// 	}

// 	encryptedPassword, err := lib.GenerateHashFromString(payload.Password)
// 	if err != nil {
// 		logger.ZapGorm.Error(ctx, "Error EncryptPassword", map[string]interface{}{
// 			"error": err,
// 			"tags":  []string{"auth", "login"},
// 		})

// 		return user, lib.ErrorInternalServer
// 	}

// 	user = model.User{
// 		Code:              fmt.Sprint("USR-", time.Now().UnixMilli()),
// 		Name:              payload.Name,
// 		Email:             payload.Email,
// 		EncryptedPassword: encryptedPassword,
// 		IsActiveFlag:      "Y",
// 		CreatedAt:         time.Now(),
// 		CreatedBy:         payload.CreatedBy,
// 		UpdatedAt:         time.Now(),
// 		UpdatedBy:         payload.CreatedBy,
// 	}

// 	err = usecase.repo.Transaction(ctx, func(ctx context.Context) error {
// 		user, err = usecase.repo.CreateUser(ctx, user)
// 		if err != nil {
// 			return err
// 		}

// 		userRole := model.UserRole{
// 			UserID:    user.ID,
// 			RoleID:    constant.RoleIDPetani,
// 			CreatedBy: user.ID,
// 			UpdatedBy: user.ID,
// 		}
// 		_, err = usecase.repo.CreateUserRole(ctx, userRole)
// 		if err != nil {
// 			return err
// 		}

// 		userPhone := model.UserPhone{
// 			UserID:       user.ID,
// 			PhoneNumber:  payload.Phone,
// 			IsActiveFlag: "Y",
// 			CreatedBy:    user.ID,
// 			UpdatedBy:    user.ID,
// 		}
// 		_, err = usecase.repo.CreateUserPhone(ctx, userPhone)
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	})

// 	return user, err
// }

func (usecase *Usecase) Login(ctx context.Context, payload request.LoginRequest) (resp response.LoginResponse, err error) {
	user := model.User{}
	isEmail := lib.IsEmail(payload.EmailOrPhone)
	if isEmail {
		user, err = usecase.repo.GetUserByEmail(ctx, payload.EmailOrPhone)
		if err != nil {
			return resp, err
		}
	} else {
		phoneUser, err := usecase.repo.GetUserPhone(ctx, payload.EmailOrPhone)
		if err != nil {
			return resp, err
		}

		user, err = usecase.repo.GetUserByID(ctx, phoneUser.UserID)
		if err != nil {
			return resp, err
		}
	}

	isValidPassword := lib.ComparePassword(user.EncryptedPassword, []byte(payload.Password))
	if !isValidPassword {
		return resp, lib.ErrorWrongEmailOrPassword
	}

	userRoles, err := usecase.repo.GetUserRoleByUserID(ctx, user.ID)
	if err != nil {
		return
	}

	role, err := usecase.repo.GetRole(ctx, request.GetRoleQuery{Role: model.Role{ID: userRoles.RoleID}})
	if err != nil {
		return
	}

	resp.ID = user.ID
	resp.Email = user.Email
	resp.Name = user.Name
	resp.PhotoURL = user.PhotoURL
	resp.ChangePasswordAt = user.ChangePasswordAt
	resp.TncAcceptedAt = user.TncAcceptedAt
	resp.Role = role

	session := model.Session{
		UserID:                   user.ID,
		Token:                    fmt.Sprint(user.ID, uuid.NewString()),
		IsLoginWithBiometricFlag: "N",
		ExpiredAt:                time.Now().AddDate(0, 1, 0), //1 month
	}
	_, err = usecase.repo.CreateSession(ctx, session)
	if err != nil {
		return
	}

	resp.Token = session.Token

	return
}
