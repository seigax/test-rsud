package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/logger"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/response"
)

func (usecase *Usecase) CheckOTPLogin(ctx context.Context, payload request.OTPRequest) (resp response.LoginResponse, err error) {
	var user model.User
	isEmail := lib.IsEmail(payload.EmailOrPhone)
	if isEmail {
		getUser, err := usecase.repo.GetUserByEmail(ctx, payload.EmailOrPhone)
		if err != nil {
			logger.Error(ctx, "Error GetUserByEmail", map[string]interface{}{
				"error": err,
				"tags":  []string{"auth", "login"},
			})
			return resp, err
		}
		user = getUser
	} else {
		phoneUser, err := usecase.repo.GetUserPhone(ctx, payload.EmailOrPhone)
		if err != nil {
			logger.Error(ctx, "Error GetUserPhone", map[string]interface{}{
				"error": err,
				"tags":  []string{"auth", "login"},
			})
			return resp, err
		}

		getUser, err := usecase.repo.GetUserByID(ctx, phoneUser.UserID)
		if err != nil {
			logger.Error(ctx, "Error GetUserByID", map[string]interface{}{
				"error": err,
				"tags":  []string{"auth", "login"},
			})
			return resp, err
		}
		user = getUser
	}

	// cek otpnya
	userVerification, err := usecase.repo.GetUserVerificationByUserId(ctx, user.ID)
	if err != nil {
		logger.Error(ctx, "Error GetUserVerificationByUserId", map[string]interface{}{
			"error": err,
			"tags":  []string{"auth", "login"},
		})
		return resp, err
	}

	if userVerification.VerificationCode != payload.Otp {
		return resp, lib.NotValidOTP
	}

	//update verification status
	err = usecase.repo.UpdateVerificationStatusToTrue(ctx, userVerification.ID)
	if err != nil {
		logger.Error(ctx, "Error UpdateVerificationStatusToTrue", map[string]interface{}{
			"error": err,
			"tags":  []string{"auth", "login"},
		})
		return resp, err
	}

	session := model.Session{
		UserID:                   user.ID,
		Token:                    fmt.Sprint(user.ID, uuid.NewString()),
		IsLoginWithBiometricFlag: "N",
		ExpiredAt:                time.Now().AddDate(0, 1, 0), //1 month
	}
	_, err = usecase.repo.CreateSession(ctx, session)
	if err != nil {
		logger.Error(ctx, "Error CreateSession", map[string]interface{}{
			"error": err,
			"tags":  []string{"auth", "login"},
		})
		return
	}

	userRoles, err := usecase.repo.GetUserRoleByUserID(ctx, user.ID)
	if err != nil {
		logger.Error(ctx, "Error GetUserRoleByUserID", map[string]interface{}{
			"error": err,
			"tags":  []string{"auth", "login"},
		})
		return
	}

	role, err := usecase.repo.GetRole(ctx, request.GetRoleQuery{Role: model.Role{ID: userRoles.RoleID}})
	if err != nil {
		logger.Error(ctx, "Error GetRoleByID", map[string]interface{}{
			"error": err,
			"tags":  []string{"auth", "login"},
		})
		return
	}

	// update verification status to true

	resp.ID = user.ID
	resp.Email = user.Email
	resp.Name = user.Name
	resp.PhotoURL = user.PhotoURL
	resp.ChangePasswordAt = user.ChangePasswordAt
	resp.TncAcceptedAt = user.TncAcceptedAt
	resp.Role = role

	resp.Token = session.Token

	return resp, nil
}

func (usecase *Usecase) ChangePassword(ctx context.Context, payload request.ChangePassword) (err error) {
	if payload.Password != payload.ConfirmationPassword {
		return lib.PasswordNotSame
	}

	encryptedPassword, err := lib.GenerateHashFromString(payload.Password)
	if err != nil {
		logger.Error(ctx, "Error GenerateHashFromString", map[string]interface{}{
			"error": err,
			"tags":  []string{"auth", "login"},
		})

		return lib.ErrorInternalServer
	}

	user, err := usecase.repo.GetUserByID(ctx, payload.UserId)
	if err != nil {
		logger.Error(ctx, "Error GetUserByID", map[string]interface{}{
			"error": err,
			"tags":  []string{"auth", "login"},
		})

		return err
	}

	user.EncryptedPassword = encryptedPassword

	_, err = usecase.repo.UpdateUser(ctx, user)
	if err != nil {
		logger.Error(ctx, "Error UpdateUser", map[string]interface{}{
			"error": err,
			"tags":  []string{"auth", "login"},
		})
		return err
	}

	return nil
}
