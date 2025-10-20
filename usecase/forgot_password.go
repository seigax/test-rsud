package usecase

import (
	"context"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/constant"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/response"
	"time"
)

func (usecase *Usecase) GetFarmerByPhone(ctx context.Context, payload request.PhoneRequest) (resp response.UserFarmerResponse, err error) {

	phoneUser, err := usecase.repo.GetUserPhone(ctx, payload.Phone)
	if err != nil {
		return resp, err
	}

	user, err := usecase.repo.GetUserByID(ctx, phoneUser.UserID)
	if err != nil {
		return resp, err
	}

	resp.ID = user.ID
	resp.Name = user.Name
	resp.Email = user.Email
	resp.Phone = phoneUser.PhoneNumber
	resp.PhotoURL = user.PhotoURL
	// disini harusnya send otp

	// create verification
	userVerification := model.UserVerification{
		VerificationCode:            lib.GenerateOTP(4),
		UserID:                      int(user.ID),
		VerificationTypeID:          constant.VerificationTypeIDOTPResetPassword,
		CommunicationDeviceTypeCode: "",
		VerificationStatusFlag:      false,
		ExpiredAt:                   lib.GenerateFutureTimeSeconds(120),
		CreatedAt:                   time.Now(),
		CreatedBy:                   int(user.ID),
		UpdatedAt:                   time.Now(),
		UpdatedBy:                   int(user.ID),
	}
	_, err = usecase.repo.CreateUserVerification(ctx, userVerification)
	if err != nil {
		return
	}

	return resp, nil
}
