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

func TestCreateRole(t *testing.T) {
	ctx := context.Background()

	Convey("CreateRole Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("Transaction", mock.Anything).Return(nil)
			r2 := m.On("CreateRole", mock.Anything).Return(model.Role{}, nil)
			r3 := m.On("UpdateRole", mock.Anything).Return(model.Role{}, nil)

			res, err := u.CreateRole(ctx, request.CreateRoleRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
		Convey("fail CreateRole", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("CreateRole", mock.Anything).Return(model.Role{}, errors.New("error"))

			res, err := u.CreateRole(ctx, request.CreateRoleRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail UpdateRole", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("CreateRole", mock.Anything).Return(model.Role{}, nil)
			r3 := m.On("UpdateRole", mock.Anything).Return(model.Role{}, errors.New("error"))

			res, err := u.CreateRole(ctx, request.CreateRoleRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
	})
}

func TestGetRoles(t *testing.T) {
	ctx := context.Background()
	query := request.GetRoleQuery{}

	Convey("GetRoles Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetRoles", query).Return([]model.Role{}, nil)
			r2 := m.On("GetRoleTotal", query).Return(uint(1), nil)

			res, err := u.GetRoles(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetRoles", func() {
			r1 := m.On("GetRoles", query).Return([]model.Role{}, errors.New("Error GetRoles"))
			r2 := m.On("GetRoleTotal", query).Return(uint(0), nil)

			res, err := u.GetRoles(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetRoleTotal", func() {
			r1 := m.On("GetRoles", query).Return([]model.Role{}, nil)
			r2 := m.On("GetRoleTotal", query).Return(uint(0), errors.New("Error GetRoleTotal"))

			res, err := u.GetRoles(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestUpdateRole(t *testing.T) {
	ctx := context.Background()

	Convey("UpdateRole Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetRoleByID", mock.Anything).Return(model.Role{ID: 1}, nil)
			r2 := m.On("UpdateRole", mock.Anything).Return(model.Role{ID: 1}, nil)

			res, err := u.UpdateRole(ctx, request.UpdateRoleRequest{ID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail UpdateRole", func() {
			r1 := m.On("GetRoleByID", mock.Anything).Return(model.Role{}, nil)
			r2 := m.On("UpdateRole", mock.Anything).Return(model.Role{}, errors.New("Error UpdateRole"))

			res, err := u.UpdateRole(ctx, request.UpdateRoleRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetRoleByID", func() {
			r1 := m.On("GetRoleByID", mock.Anything).Return(model.Role{}, errors.New("ERROR"))
			r2 := m.On("UpdateRole", mock.Anything).Return(model.Role{}, nil)

			res, err := u.UpdateRole(ctx, request.UpdateRoleRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestGetRoleByID(t *testing.T) {

	ctx := context.Background()

	Convey("GetRoleDetail Usecase Scenario", t, func() {
		Convey("success", func() {
			r := m.On("GetRoleByID", uint(1)).Return(model.Role{ID: 1}, nil)

			res, err := u.GetRoleDetail(ctx, uint(1))

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r.Unset()
		})
		Convey("fail", func() {
			r := m.On("GetRoleByID", uint(0)).Return(model.Role{ID: 0}, errors.New("Role Not Found"))

			res, err := u.GetRoleDetail(ctx, uint(0))

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r.Unset()
		})
	})
}

func TestDeleteRole(t *testing.T) {

	ctx := context.Background()

	Convey("DeleteRole Usecase Scenario", t, func() {
		Convey("success", func() {
			r := m.On("DeleteRole", uint(1)).Return(nil)

			err := u.DeleteRole(ctx, uint(1))

			So(err, ShouldBeNil)

			r.Unset()
		})
		Convey("fail", func() {
			r := m.On("DeleteRole", uint(0)).Return(errors.New("Role Not Found"))

			err := u.DeleteRole(ctx, uint(0))

			So(err, ShouldNotBeNil)

			r.Unset()
		})
	})
}
