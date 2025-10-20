package usecase

import (
	"context"
	"time"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
)

func (usecase *Usecase) ValidateToken(ctx context.Context, token string) (session model.Session, err error) {
	session, err = usecase.repo.GetSession(ctx, model.Session{
		Token: token,
	})
	if err != nil {
		return
	}

	if session.IsExpired() {
		err = lib.ErrorExpiredToken
		return
	}

	return
}

func (usecase *Usecase) ExpireToken(ctx context.Context, session model.Session) error {
	session.ExpiredAt = time.Now()
	_, err := usecase.repo.UpdateSession(ctx, session)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *Usecase) AuthorizeUserForRole(ctx context.Context, currentUserID, targetRoleID uint) (bool, error) {
	userRole, err := usecase.repo.GetUserRoleByUserID(ctx, currentUserID)
	if err != nil {
		return false, lib.ErrorForbidden
	}

	if userRole.RoleID != targetRoleID {
		return false, lib.ErrorForbidden
	}

	return true, nil
}

func (usecase *Usecase) AuthorizeUserWithRolesForMenu(ctx context.Context, userRoles []model.UserRole, feUrls []string) (bool, error) {
	for _, userRole := range userRoles {
		roleMenus, err := usecase.repo.GetRoleMenusWithRoleUrlByRoleID(ctx, userRole.RoleID)
		if err != nil {
			return false, lib.ErrorForbidden
		}

		for _, roleMenu := range roleMenus {
			for _, feUrl := range feUrls {
				if roleMenu.MenuUrl == feUrl {
					return true, nil
				}
			}
		}
	}

	return false, nil
}
