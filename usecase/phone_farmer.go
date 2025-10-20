package usecase

import (
	"context"
	"errors"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
	"time"
)

func (usecase *Usecase) GetListUserPhone(ctx context.Context, userId uint) ([]model.UserPhone, error) {
	listphone, err := usecase.repo.GetAllUserPhones(ctx, userId)
	if err != nil {
		return nil, err
	}

	return listphone, nil
}

func (usecase *Usecase) AddUserPhone(ctx context.Context, payload request.ReqAddPhone) (*model.UserPhone, error) {
	// find user phone by user_id
	_, err := usecase.repo.GetUserPhone(ctx, payload.Phone)
	if err != nil {
		if errors.Is(err, lib.ErrorNotFound) {
			userPhone := model.UserPhone{
				UserID:       payload.UserId,
				PhoneNumber:  payload.Phone,
				IsActiveFlag: "N",
				CreatedAt:    time.Now(),
				CreatedBy:    payload.UserId,
				UpdatedAt:    time.Now(),
				UpdatedBy:    payload.UserId,
			}

			phone, err := usecase.repo.CreateUserPhone(ctx, userPhone)
			if err != nil {
				return nil, err
			}

			return &phone, nil
		}
		// handle other errors
		return nil, err
	}

	return nil, lib.ErrorPhoneAlreadyRegistered
}

func (usecase *Usecase) DeleteUserPhone(ctx context.Context, id uint) error {
	err := usecase.repo.DeleteUserPhone(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *Usecase) UpdateUserPhone(ctx context.Context, payload request.ReqUpdatePhone) (*model.UserPhone, error) {
	// find user phone by user_id
	userPhone, err := usecase.repo.GetUserPhoneByID(ctx, payload.Id)
	if err != nil {
		return nil, err
	}

	_, err = usecase.repo.GetUserPhone(ctx, payload.Phone)
	if err != nil {
		if errors.Is(err, lib.ErrorNotFound) {
			userPhone.PhoneNumber = payload.Phone
			userPhone.UpdatedBy = payload.UserId
			userPhone.UpdatedAt = time.Now()

			phone, err := usecase.repo.UpdateUserPhone(ctx, userPhone)
			if err != nil {
				return nil, err
			}

			return &phone, nil
		}
		// handle other errors
		return nil, err
	}

	return nil, lib.ErrorPhoneAlreadyRegistered

}

func (usecase *Usecase) ChangeActiveUserPhone(ctx context.Context, payload request.ReqUpdatePhone) error {
	err := usecase.repo.ChangeActiveUserPhone(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}
