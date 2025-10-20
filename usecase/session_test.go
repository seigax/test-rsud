package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/constant"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
)

func TestValidateToken(t *testing.T) {
	ctx := context.Background()

	Convey("ValidateToken Usecase Scenario", t, func() {
		Convey("success", func() {
			r := m.On("GetSessionByToken", mock.Anything).Return(model.Session{ExpiredAt: time.Now().AddDate(0, 1, 0)}, nil)

			res, err := u.ValidateToken(ctx, "")

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r.Unset()
		})
		Convey("fail", func() {
			r := m.On("GetSessionByToken", mock.Anything).Return(model.Session{}, errors.New("Error ValidateToken"))

			res, err := u.ValidateToken(ctx, "")

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r.Unset()
		})
		Convey("fail on expired token", func() {
			r := m.On("GetSessionByToken", mock.Anything).Return(model.Session{ExpiredAt: time.Now().AddDate(0, 0, -1)}, nil)

			res, err := u.ValidateToken(ctx, "")

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r.Unset()
		})
	})
}

func TestAuthorizeUserForRole(t *testing.T) {
	ctx := context.Background()

	Convey("AuthorizeUserForRole Usecase Scenario", t, func() {
		Convey("success on admin", func() {
			r := m.On("GetUserRoleByUserID", mock.Anything).Return(model.UserRole{RoleID: constant.RoleIDAdministrator}, nil)

			res, err := u.AuthorizeUserForRole(ctx, 1, constant.RoleIDAdministrator)

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r.Unset()
		})
		Convey("fail on error", func() {
			r := m.On("GetUserRoleByUserID", mock.Anything).Return(model.UserRole{}, errors.New("ERROR"))

			res, err := u.AuthorizeUserForRole(ctx, 1, constant.RoleIDAdministrator)

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r.Unset()
		})
	})
}
