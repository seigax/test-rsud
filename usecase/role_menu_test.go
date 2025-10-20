package usecase

import (
	"context"
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

func TestSaveRoleMenu(t *testing.T) {
	ctx := context.Background()

	Convey("SaveRoleMenu Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("Transaction", mock.Anything).Return(nil)
			r2 := m.On("DeleteAllRoleMenu", uint(1)).Return(nil)
			r3 := m.On("CreateRoleMenu", mock.Anything).Return(model.RoleMenu{}, nil)

			err := u.SaveRoleMenu(ctx, request.CreateRoleMenuRequest{RoleID: 1, Menu: []uint{1, 2}})

			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
		Convey("fail DeleteAllRoleMenu", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("DeleteAllRoleMenu", uint(1)).Return(errors.New("error"))
			r3 := m.On("CreateRoleMenu", mock.Anything).Return(model.RoleMenu{}, nil)

			err := u.SaveRoleMenu(ctx, request.CreateRoleMenuRequest{RoleID: 1, Menu: []uint{1, 2}})

			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
		Convey("fail CreateRoleMenu", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("DeleteAllRoleMenu", uint(1)).Return(nil)
			r3 := m.On("CreateRoleMenu", mock.Anything).Return(model.RoleMenu{}, errors.New("error"))

			err := u.SaveRoleMenu(ctx, request.CreateRoleMenuRequest{RoleID: 1, Menu: []uint{1, 2}})

			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
	})
}
