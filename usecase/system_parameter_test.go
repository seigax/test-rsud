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

func TestCreateSystemParameter(t *testing.T) {
	ctx := context.Background()

	Convey("CreateSystemParameter Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("Transaction", mock.Anything).Return(nil)
			r2 := m.On("CreateSystemParameter", mock.Anything).Return(model.SystemParameter{}, nil)
			r3 := m.On("UpdateSystemParameter", mock.Anything).Return(model.SystemParameter{}, nil)

			res, err := u.CreateSystemParameter(ctx, request.CreateSystemParameterRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
		Convey("fail", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("CreateSystemParameter", mock.Anything).Return(model.SystemParameter{}, errors.New("Error CreateSystemParameter"))

			res, err := u.CreateSystemParameter(ctx, request.CreateSystemParameterRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestGetSystemParameters(t *testing.T) {
	ctx := context.Background()
	query := request.GetSystemParameterQuery{}

	Convey("GetSystemParameters Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetSystemParameters", query).Return([]model.SystemParameter{}, nil)
			r2 := m.On("GetSystemParameterTotal", query).Return(uint(2), nil)

			res, err := u.GetSystemParameters(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetSystemParameters", func() {
			r1 := m.On("GetSystemParameters", query).Return([]model.SystemParameter{}, errors.New("Error GetSystemParameters"))
			r2 := m.On("GetSystemParameterTotal", query).Return(uint(0), errors.New("Error GetSystemParameterTotal"))

			res, err := u.GetSystemParameters(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetSystemParameterTotal", func() {
			r1 := m.On("GetSystemParameters", query).Return([]model.SystemParameter{}, nil)
			r2 := m.On("GetSystemParameterTotal", query).Return(uint(0), errors.New("Error GetSystemParameterTotal"))

			res, err := u.GetSystemParameters(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestUpdateSystemParameter(t *testing.T) {
	ctx := context.Background()

	Convey("UpdateSystemParameter Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetSystemParameterByID", mock.Anything).Return(model.SystemParameter{}, nil)
			r2 := m.On("UpdateSystemParameter", mock.Anything).Return(model.SystemParameter{}, nil)

			res, err := u.UpdateSystemParameter(ctx, request.UpdateSystemParameterRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail UpdateSystemParameter", func() {
			r1 := m.On("GetSystemParameterByID", mock.Anything).Return(model.SystemParameter{ID: 1}, nil)
			r2 := m.On("UpdateSystemParameter", mock.Anything).Return(model.SystemParameter{}, errors.New("Error UpdateSystemParameter"))

			res, err := u.UpdateSystemParameter(ctx, request.UpdateSystemParameterRequest{ID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetSystemParameterByID", func() {
			r1 := m.On("GetSystemParameterByID", mock.Anything).Return(model.SystemParameter{ID: 1}, errors.New("error"))
			r2 := m.On("UpdateSystemParameter", mock.Anything).Return(model.SystemParameter{}, errors.New("Error UpdateSystemParameter"))

			res, err := u.UpdateSystemParameter(ctx, request.UpdateSystemParameterRequest{ID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestGetSystemParameterByID(t *testing.T) {

	ctx := context.Background()

	Convey("GetSystemParameterDetail Usecase Scenario", t, func() {
		Convey("success", func() {
			r := m.On("GetSystemParameterByID", uint(1)).Return(model.SystemParameter{ID: 1}, nil)

			res, err := u.GetSystemParameterDetail(ctx, uint(1))

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r.Unset()
		})
		Convey("fail", func() {
			r := m.On("GetSystemParameterByID", uint(0)).Return(model.SystemParameter{ID: 0}, errors.New("SystemParameter Not Found"))

			res, err := u.GetSystemParameterDetail(ctx, uint(0))

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r.Unset()
		})
	})
}

func TestDeleteSystemParameter(t *testing.T) {

	ctx := context.Background()

	Convey("DeleteSystemParameter Usecase Scenario", t, func() {
		Convey("success", func() {
			r := m.On("DeleteSystemParameter", uint(1)).Return(nil)

			err := u.DeleteSystemParameter(ctx, uint(1))

			So(err, ShouldBeNil)

			r.Unset()
		})
		Convey("fail", func() {
			r := m.On("DeleteSystemParameter", uint(0)).Return(errors.New("SystemParameter Not Found"))

			err := u.DeleteSystemParameter(ctx, uint(0))

			So(err, ShouldNotBeNil)

			r.Unset()
		})
	})
}
