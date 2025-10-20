package usecase

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/response"
)

func (usecase *Usecase) CreateRole(ctx context.Context, payload request.CreateRoleRequest) (role model.Role, err error) {
	err = usecase.repo.Transaction(ctx, func(ctx context.Context) error {
		role = model.Role{
			Name:         payload.Name,
			Description:  payload.Description,
			Type:         payload.Type,
			Platform:     payload.Platform,
			IsActiveFlag: payload.IsActiveFlag,
			CreatedAt:    time.Now(),
			CreatedBy:    payload.CreatedBy,
			UpdatedAt:    time.Now(),
			UpdatedBy:    payload.CreatedBy,
		}

		role, err = usecase.repo.CreateRole(ctx, role)
		if err != nil {
			return err
		}

		role.Code = fmt.Sprintf("ROLE-%04d", role.ID)

		role, err = usecase.repo.UpdateRole(ctx, role)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

func (usecase *Usecase) GetRoles(ctx context.Context, query request.GetRoleQuery) (response.GetRoleResponse, error) {
	var res response.GetRoleResponse

	roles, err := usecase.repo.GetRoles(ctx, query)
	if err != nil {
		return res, err
	}

	total, err := usecase.repo.GetRoleTotal(ctx, query)
	if err != nil {
		return res, err
	}

	res.Roles = roles
	res.Total = total
	res.Limit = query.Limit
	res.Offset = query.GetOffset()
	res.Page = query.Page

	return res, nil
}

func (usecase *Usecase) GetRoleDetail(ctx context.Context, ID uint) (model.Role, error) {
	role, err := usecase.repo.GetRole(ctx, request.GetRoleQuery{Role: model.Role{ID: ID}})
	if err != nil {
		return role, err
	}

	return role, nil
}

func (usecase *Usecase) UpdateRole(ctx context.Context, payload request.UpdateRoleRequest) (model.Role, error) {
	role, err := usecase.repo.GetRole(ctx, request.GetRoleQuery{Role: model.Role{ID: payload.ID}})
	if err != nil {
		return role, err
	}

	role.Name = payload.Name
	role.Description = payload.Description
	role.Type = payload.Type
	role.Platform = payload.Platform
	role.IsActiveFlag = payload.IsActiveFlag
	role.UpdatedAt = time.Now()
	role.UpdatedBy = payload.UpdatedBy

	res, err := usecase.repo.UpdateRole(ctx, role)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (usecase *Usecase) DeleteRole(ctx context.Context, ID uint) error {
	err := usecase.repo.DeleteRole(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
