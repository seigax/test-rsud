package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/logger"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/response"
)

func (usecase *Usecase) CreateUser(ctx context.Context, payload request.CreateUserRequest) (model.User, error) {
	var user model.User
	err := usecase.repo.Transaction(ctx, func(ctx context.Context) error {

		encryptedPassword, err := lib.GenerateHashFromString(payload.Password)
		if err != nil {
			logger.Error(ctx, "Error EncryptPassword", map[string]interface{}{
				"error": err,
				"tags":  []string{"auth", "login"},
			})

			return lib.ErrorInternalServer
		}

		user = model.User{
			Name:              payload.Name,
			Email:             payload.Email,
			EncryptedPassword: encryptedPassword,
			IsActiveFlag:      payload.IsActiveFlag,
			CreatedAt:         time.Now(),
			CreatedBy:         payload.CreatedBy,
			UpdatedAt:         time.Now(),
			UpdatedBy:         payload.CreatedBy,
		}

		user, err = usecase.repo.CreateUser(ctx, user)
		if err != nil {
			return err
		}

		user.Code = fmt.Sprintf("USR-%013d", user.ID)

		_, err = usecase.repo.UpdateUser(ctx, user)
		if err != nil {
			return err
		}

		for _, roleID := range payload.Roles {
			_, err = usecase.repo.CreateUserRole(ctx, model.UserRole{
				UserID: user.ID,
				RoleID: roleID,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})

	return user, err
}

func (usecase *Usecase) GetUsers(ctx context.Context, query request.GetUserQuery) (response.GetUserResponse, error) {
	var res response.GetUserResponse

	users, err := usecase.repo.GetUsers(ctx, query)
	if err != nil {
		return res, err
	}

	total, err := usecase.repo.GetUserTotal(ctx, query)
	if err != nil {
		return res, err
	}

	for _, user := range users {
		roles, _ := usecase.repo.GetUserRolesWithRoleDetailByUserID(ctx, user.ID)

		roleStrings := []string{}
		roleType := ""

		for _, role := range roles {
			roleStrings = append(roleStrings, role.RoleName)
			roleType = role.RoleType
		}

		res.Users = append(res.Users, response.GetUserWithRolesResponse{
			User:     user,
			Roles:    strings.Join(roleStrings, ", "),
			RoleType: roleType,
		})
	}

	res.Total = total
	res.Limit = query.Limit
	res.Offset = query.GetOffset()
	res.Page = query.Page

	return res, nil
}

func (usecase *Usecase) GetUserDetail(ctx context.Context, ID uint) (model.User, error) {
	var user model.User

	user, err := usecase.repo.GetUserByID(ctx, ID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (usecase *Usecase) UpdateUser(ctx context.Context, payload request.UpdateUserRequest) (user model.User, err error) {
	err = usecase.repo.Transaction(ctx, func(ctx context.Context) error {
		user, err = usecase.repo.GetUserByID(ctx, payload.UserID)
		if err != nil {
			return err
		}

		if payload.Password != "" {
			encryptedPassword, err := lib.GenerateHashFromString(payload.Password)
			if err != nil {
				logger.Error(ctx, "Error EncryptPassword", map[string]interface{}{
					"error": err,
					"tags":  []string{"auth", "login"},
				})

				return lib.ErrorInternalServer
			}

			user.EncryptedPassword = encryptedPassword
		}

		user.Name = payload.Name
		user.Email = payload.Email
		user.IsActiveFlag = payload.IsActiveFlag
		user.UpdatedAt = time.Now()
		user.UpdatedBy = payload.UpdatedBy

		user, err = usecase.repo.UpdateUser(ctx, user)
		if err != nil {
			return err
		}

		err = usecase.repo.DeleteUserRoleByUserID(ctx, user.ID)
		if err != nil {
			return err
		}

		for _, roleID := range payload.Roles {
			_, err = usecase.repo.CreateUserRole(ctx, model.UserRole{
				UserID: user.ID,
				RoleID: roleID,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})

	return user, err
}

func (usecase *Usecase) DeleteUser(ctx context.Context, ID uint) error {
	err := usecase.repo.DeleteUser(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
