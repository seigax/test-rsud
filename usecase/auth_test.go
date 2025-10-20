package usecase

import (
	"context"
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/mock"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/constant"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
)

// func TestRegisterFarmer(t *testing.T) {
// 	ctx := context.Background()

// 	Convey("RegisterFarmer Usecase Scenario", t, func() {
// 		Convey("success", func() {
// 			r1 := m.On("GetUserByEmail", mock.Anything).Return(model.User{}, lib.ErrorNotFound)
// 			r2 := m.On("CheckPhoneNumberIsUsed", mock.Anything).Return(false, lib.ErrorNotFound)
// 			r3 := m.On("Transaction", mock.Anything).Return(nil)
// 			r4 := m.On("CreateUser", mock.Anything).Return(model.User{}, nil)
// 			r5 := m.On("CreateUserRole", mock.Anything).Return(model.UserRole{}, nil)
// 			r6 := m.On("CreateUserPhone", mock.Anything).Return(model.UserPhone{}, nil)

// 			res, err := u.RegisterFarmer(ctx, request.RegisterFarmerRequest{})

// 			So(res, ShouldNotBeNil)
// 			So(err, ShouldBeNil)

// 			r1.Unset()
// 			r2.Unset()
// 			r3.Unset()
// 			r4.Unset()
// 			r5.Unset()
// 			r6.Unset()
// 		})
// 		Convey("fail on CreateUser repo", func() {
// 			r1 := m.On("GetUserByEmail", mock.Anything).Return(model.User{}, lib.ErrorNotFound)
// 			r2 := m.On("CheckPhoneNumberIsUsed", mock.Anything).Return(false, lib.ErrorNotFound)
// 			r3 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
// 			r4 := m.On("CreateUser", mock.Anything).Return(model.User{}, errors.New("error"))
// 			r5 := m.On("CreateUserRole", mock.Anything).Return(model.UserRole{}, nil)
// 			r6 := m.On("CreateUserPhone", mock.Anything).Return(model.UserPhone{}, nil)

// 			res, err := u.RegisterFarmer(ctx, request.RegisterFarmerRequest{})

// 			So(res, ShouldNotBeNil)
// 			So(err, ShouldNotBeNil)

// 			r1.Unset()
// 			r2.Unset()
// 			r3.Unset()
// 			r4.Unset()
// 			r5.Unset()
// 			r6.Unset()
// 		})
// 		Convey("fail on email already used", func() {
// 			r1 := m.On("GetUserByEmail", mock.Anything).Return(model.User{ID: 1}, nil)
// 			r2 := m.On("CheckPhoneNumberIsUsed", mock.Anything).Return(false, nil)
// 			r3 := m.On("Transaction", mock.Anything).Return(nil)

// 			res, err := u.RegisterFarmer(ctx, request.RegisterFarmerRequest{})

// 			So(res, ShouldNotBeNil)
// 			So(err, ShouldNotBeNil)

// 			r1.Unset()
// 			r2.Unset()
// 			r3.Unset()
// 		})
// 		Convey("fail on GetUserByEmail db error", func() {
// 			r1 := m.On("GetUserByEmail", mock.Anything).Return(model.User{}, errors.New("DB ERROR"))
// 			r2 := m.On("CheckPhoneNumberIsUsed", mock.Anything).Return(false, nil)
// 			r3 := m.On("Transaction", mock.Anything).Return(nil)

// 			res, err := u.RegisterFarmer(ctx, request.RegisterFarmerRequest{})

// 			So(res, ShouldNotBeNil)
// 			So(err, ShouldNotBeNil)

// 			r1.Unset()
// 			r2.Unset()
// 			r3.Unset()
// 		})
// 		Convey("fail on phone already used", func() {
// 			r1 := m.On("GetUserByEmail", mock.Anything).Return(model.User{}, nil)
// 			r2 := m.On("CheckPhoneNumberIsUsed", mock.Anything).Return(true, nil)
// 			r3 := m.On("Transaction", mock.Anything).Return(nil)

// 			res, err := u.RegisterFarmer(ctx, request.RegisterFarmerRequest{})

// 			So(res, ShouldNotBeNil)
// 			So(err, ShouldNotBeNil)

// 			r1.Unset()
// 			r2.Unset()
// 			r3.Unset()
// 		})
// 		Convey("fail on CheckPhoneNumberIsUsed db error", func() {
// 			r1 := m.On("GetUserByEmail", mock.Anything).Return(model.User{}, nil)
// 			r2 := m.On("CheckPhoneNumberIsUsed", mock.Anything).Return(false, errors.New("DB ERROR"))
// 			r3 := m.On("Transaction", mock.Anything).Return(nil)

// 			res, err := u.RegisterFarmer(ctx, request.RegisterFarmerRequest{})

// 			So(res, ShouldNotBeNil)
// 			So(err, ShouldNotBeNil)

// 			r1.Unset()
// 			r2.Unset()
// 			r3.Unset()
// 		})
// 	})
// }

func TestLogin(t *testing.T) {
	ctx := context.Background()
	passwordHash, _ := lib.GenerateHashFromString("123456")

	Convey("Login Usecase Scenario", t, func() {
		Convey("success on email", func() {
			r1 := m.On("GetUserByEmail", mock.Anything).Return(model.User{EncryptedPassword: passwordHash}, nil)
			r2 := m.On("GetUserRoleByUserID", mock.Anything).Return(model.UserRole{}, nil)
			r3 := m.On("GetRoleByID", mock.Anything).Return(model.Role{ID: constant.RoleIDAdministrator}, nil)
			r4 := m.On("CreateSession", mock.Anything).Return(model.Session{}, nil)

			res, err := u.Login(ctx, request.LoginRequest{EmailOrPhone: "a@a.com", Password: "123456"})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
			r4.Unset()
		})
		Convey("success on phone", func() {
			r1 := m.On("GetUserPhone", mock.Anything).Return(model.UserPhone{}, nil)
			r2 := m.On("GetUserByID", mock.Anything).Return(model.User{EncryptedPassword: passwordHash}, nil)
			r3 := m.On("GetUserRoleByUserID", mock.Anything).Return(model.UserRole{}, nil)
			r4 := m.On("GetRoleByID", mock.Anything).Return(model.Role{ID: constant.RoleIDAdministrator}, nil)
			r5 := m.On("CreateSession", mock.Anything).Return(model.Session{}, nil)

			res, err := u.Login(ctx, request.LoginRequest{EmailOrPhone: "0813111111", Password: "123456"})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
			r4.Unset()
			r5.Unset()
		})
		Convey("fail on email", func() {
			r1 := m.On("GetUserByEmail", mock.Anything).Return(model.User{}, errors.New("error GetUserByEmail"))
			r2 := m.On("GetUserRoleByUserID", mock.Anything).Return(model.UserRole{}, nil)
			r3 := m.On("GetRoleByID", mock.Anything).Return(model.Role{}, nil)
			r4 := m.On("CreateSession", mock.Anything).Return(model.Session{}, nil)

			res, err := u.Login(ctx, request.LoginRequest{EmailOrPhone: "a@a.com"})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
			r4.Unset()
		})
		Convey("fail on phone", func() {
			r1 := m.On("GetUserPhone", mock.Anything).Return(model.UserPhone{}, errors.New("ERROR GetUserPhone"))
			r2 := m.On("GetUserByID", mock.Anything).Return(model.User{}, nil)
			r3 := m.On("GetUserRoleByUserID", mock.Anything).Return(model.UserRole{}, nil)
			r4 := m.On("GetRoleByID", mock.Anything).Return(model.Role{}, nil)
			r5 := m.On("CreateSession", mock.Anything).Return(model.Session{}, nil)

			res, err := u.Login(ctx, request.LoginRequest{EmailOrPhone: "0811122"})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
			r4.Unset()
			r5.Unset()
		})
		Convey("fail on incorrect password", func() {
			r1 := m.On("GetUserPhone", mock.Anything).Return(model.UserPhone{}, nil)
			r2 := m.On("GetUserByID", mock.Anything).Return(model.User{EncryptedPassword: passwordHash}, nil)
			r3 := m.On("GetUserRoleByUserID", mock.Anything).Return(model.UserRole{}, nil)
			r4 := m.On("GetRoleByID", mock.Anything).Return(model.Role{}, nil)
			r5 := m.On("CreateSession", mock.Anything).Return(model.Session{}, nil)

			res, err := u.Login(ctx, request.LoginRequest{EmailOrPhone: "0811122", Password: "1234567"})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)
			So(err, ShouldResemble, lib.ErrorWrongEmailOrPassword)

			r1.Unset()
			r2.Unset()
			r3.Unset()
			r4.Unset()
			r5.Unset()
		})
	})
}
