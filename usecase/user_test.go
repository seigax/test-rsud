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

func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	Convey("CreateUser Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("Transaction", mock.Anything).Return(nil)
			r2 := m.On("CreateUser", mock.Anything).Return(model.User{}, nil)
			r3 := m.On("UpdateUser", mock.Anything).Return(model.User{}, nil)
			r4 := m.On("CreateUserRole", mock.Anything).Return(model.UserRole{}, nil)

			res, err := u.CreateUser(ctx, request.CreateUserRequest{
				Roles: []uint{1, 2},
			})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
			r4.Unset()
		})
		Convey("fail", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("CreateUser", mock.Anything).Return(model.User{}, errors.New("error"))

			res, err := u.CreateUser(ctx, request.CreateUserRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestGetUsers(t *testing.T) {
	ctx := context.Background()
	query := request.GetUserQuery{}

	Convey("GetUsers Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetUsers", query).Return([]model.User{{ID: 1}}, nil)
			r2 := m.On("GetUserTotal", query).Return(uint(1), nil)
			r3 := m.On("GetUserRolesWithRoleDetailByUserID", uint(1)).Return([]model.UserRoleWithRoleDetail{{RoleName: "Role1"}, {RoleName: "Role2"}}, nil)

			res, err := u.GetUsers(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
		Convey("fail GetUsers", func() {
			r1 := m.On("GetUsers", query).Return([]model.User{}, errors.New("Error GetUsers"))
			r2 := m.On("GetUserTotal", query).Return(uint(0), errors.New("Error GetUserTotal"))

			res, err := u.GetUsers(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetUserTotal", func() {
			r1 := m.On("GetUsers", query).Return([]model.User{}, nil)
			r2 := m.On("GetUserTotal", query).Return(uint(0), errors.New("Error GetUserTotal"))

			res, err := u.GetUsers(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestUpdateUser(t *testing.T) {
	ctx := context.Background()

	Convey("UpdateUser Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("Transaction", mock.Anything).Return(nil)
			r2 := m.On("GetUserByID", mock.Anything).Return(model.User{ID: 1}, nil)
			r3 := m.On("UpdateUser", mock.Anything).Return(model.User{ID: 1}, nil)
			r4 := m.On("DeleteUserRoleByUserID", mock.Anything).Return(nil)
			r5 := m.On("CreateUserRole", mock.Anything).Return(model.UserRole{}, nil)

			res, err := u.UpdateUser(ctx, request.UpdateUserRequest{UserID: 1, Roles: []uint{1, 2}, Password: "-"})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
			r4.Unset()
			r5.Unset()
		})
		Convey("fail GetUserByID", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("GetUserByID", mock.Anything).Return(model.User{ID: 1}, errors.New("error"))

			res, err := u.UpdateUser(ctx, request.UpdateUserRequest{UserID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail UpdateUser", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("GetUserByID", mock.Anything).Return(model.User{ID: 1}, nil)
			r3 := m.On("UpdateUser", mock.Anything).Return(model.User{}, errors.New("Error UpdateUser"))

			res, err := u.UpdateUser(ctx, request.UpdateUserRequest{UserID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
	})
}

func TestGetUserByID(t *testing.T) {

	ctx := context.Background()

	Convey("GetUserDetail Usecase Scenario", t, func() {
		Convey("success", func() {
			r := m.On("GetUserByID", uint(1)).Return(model.User{ID: 1}, nil)

			res, err := u.GetUserDetail(ctx, uint(1))

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r.Unset()
		})
		Convey("fail", func() {
			r := m.On("GetUserByID", uint(0)).Return(model.User{ID: 0}, errors.New("User Not Found"))

			res, err := u.GetUserDetail(ctx, uint(0))

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r.Unset()
		})
	})
}

func TestDeleteUser(t *testing.T) {

	ctx := context.Background()

	Convey("DeleteUser Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("Transaction", mock.Anything).Return(nil)
			r2 := m.On("DeleteUser", uint(1)).Return(nil)

			err := u.DeleteUser(ctx, uint(1))

			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("DeleteUser", uint(0)).Return(errors.New("error"))

			err := u.DeleteUser(ctx, uint(0))

			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}
