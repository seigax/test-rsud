package usecase

import (
	"context"
	"time"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

func (usecase *Usecase) SaveRoleMenu(ctx context.Context, payload request.CreateRoleMenuRequest) (err error) {
	err = usecase.repo.Transaction(ctx, func(ctx context.Context) error {
		err = usecase.repo.DeleteAllRoleMenu(ctx, payload.RoleID)

		for _, v := range payload.Menu {
			roleMenu := model.RoleMenu{
				RoleID:    payload.RoleID,
				MenuID:    v,
				CreatedAt: time.Now(),
				CreatedBy: payload.CreatedBy,
				UpdatedAt: time.Now(),
				UpdatedBy: payload.CreatedBy,
			}

			_, err = usecase.repo.CreateRoleMenu(ctx, roleMenu)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return
}
