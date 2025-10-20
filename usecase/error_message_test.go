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

func TestCreateErrorMessage(t *testing.T) {
	ctx := context.Background()

	Convey("CreateErrorMessage Usecase Scenario", t, func() {
		Convey("success", func() {
			r := m.On("CreateErrorMessage", mock.Anything).Return(model.ErrorMessage{}, nil)

			res, err := u.CreateErrorMessage(ctx, request.CreateErrorMessageRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r.Unset()
		})
		Convey("fail", func() {
			r := m.On("CreateErrorMessage", mock.Anything).Return(model.ErrorMessage{}, errors.New("Error CreateErrorMessage"))

			res, err := u.CreateErrorMessage(ctx, request.CreateErrorMessageRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r.Unset()
		})
	})
}

func TestGetErrorMessages(t *testing.T) {
	ctx := context.Background()
	query := request.GetErrorMessageQuery{}

	Convey("GetErrorMessages Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetErrorMessages", query).Return([]model.ErrorMessage{}, nil)
			r2 := m.On("GetErrorMessageTotal", query).Return(uint(2), nil)

			res, err := u.GetErrorMessages(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetErrorMessages", func() {
			r1 := m.On("GetErrorMessages", query).Return([]model.ErrorMessage{}, errors.New("Error GetErrorMessages"))
			r2 := m.On("GetErrorMessageTotal", query).Return(uint(0), errors.New("Error GetErrorMessageTotal"))

			res, err := u.GetErrorMessages(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetErrorMessageTotal", func() {
			r1 := m.On("GetErrorMessages", query).Return([]model.ErrorMessage{}, nil)
			r2 := m.On("GetErrorMessageTotal", query).Return(uint(0), errors.New("Error GetErrorMessageTotal"))

			res, err := u.GetErrorMessages(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestUpdateErrorMessage(t *testing.T) {
	ctx := context.Background()

	Convey("UpdateErrorMessage Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetErrorMessageByID", mock.Anything).Return(model.ErrorMessage{}, nil)
			r2 := m.On("UpdateErrorMessage", mock.Anything).Return(model.ErrorMessage{}, nil)

			res, err := u.UpdateErrorMessage(ctx, request.UpdateErrorMessageRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail UpdateErrorMessage", func() {
			r1 := m.On("GetErrorMessageByID", mock.Anything).Return(model.ErrorMessage{ID: 1}, nil)
			r2 := m.On("UpdateErrorMessage", mock.Anything).Return(model.ErrorMessage{}, errors.New("Error UpdateErrorMessage"))

			res, err := u.UpdateErrorMessage(ctx, request.UpdateErrorMessageRequest{ID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetErrorMessageByID", func() {
			r1 := m.On("GetErrorMessageByID", mock.Anything).Return(model.ErrorMessage{ID: 1}, errors.New("error"))
			r2 := m.On("UpdateErrorMessage", mock.Anything).Return(model.ErrorMessage{}, errors.New("Error UpdateErrorMessage"))

			res, err := u.UpdateErrorMessage(ctx, request.UpdateErrorMessageRequest{ID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestGetErrorMessageByID(t *testing.T) {

	ctx := context.Background()

	Convey("GetErrorMessageDetail Usecase Scenario", t, func() {
		Convey("success", func() {
			r := m.On("GetErrorMessageByID", uint(1)).Return(model.ErrorMessage{ID: 1}, nil)

			res, err := u.GetErrorMessageDetail(ctx, uint(1))

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r.Unset()
		})
		Convey("fail", func() {
			r := m.On("GetErrorMessageByID", uint(0)).Return(model.ErrorMessage{ID: 0}, errors.New("ErrorMessage Not Found"))

			res, err := u.GetErrorMessageDetail(ctx, uint(0))

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r.Unset()
		})
	})
}

func TestDeleteErrorMessage(t *testing.T) {

	ctx := context.Background()

	Convey("DeleteErrorMessage Usecase Scenario", t, func() {
		Convey("success", func() {
			r := m.On("DeleteErrorMessage", uint(1)).Return(nil)

			err := u.DeleteErrorMessage(ctx, uint(1))

			So(err, ShouldBeNil)

			r.Unset()
		})
		Convey("fail", func() {
			r := m.On("DeleteErrorMessage", uint(0)).Return(errors.New("ErrorMessage Not Found"))

			err := u.DeleteErrorMessage(ctx, uint(0))

			So(err, ShouldNotBeNil)

			r.Unset()
		})
	})
}

func TestGetErrorMessageByCode(t *testing.T) {

	ctx := context.Background()

	Convey("GetErrorMessageDetail Usecase Scenario", t, func() {
		Convey("success", func() {
			r := m.On("GetErrorMessageByCode", mock.Anything).Return(model.ErrorMessage{}, nil)

			res, err := u.GetErrorMessageByCode(ctx, "400")

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r.Unset()
		})
		Convey("fail", func() {
			r := m.On("GetErrorMessageByCode", mock.Anything).Return(model.ErrorMessage{}, errors.New("ErrorMessage Not Found"))

			res, err := u.GetErrorMessageByCode(ctx, "400")

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r.Unset()
		})
	})
}

func TestGetErrorMessageByCodeAndAppName(t *testing.T) {

	ctx := context.Background()

	Convey("GetErrorMessageByCodeAndAppName Usecase Scenario", t, func() {
		Convey("success", func() {
			r := m.On("GetErrorMessageByCodeAndAppName", mock.Anything).Return(model.ErrorMessage{}, nil)

			res, err := u.GetErrorMessageByCodeAndAppName(ctx, "400", "")

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r.Unset()
		})
		Convey("fail", func() {
			r := m.On("GetErrorMessageByCodeAndAppName", mock.Anything).Return(model.ErrorMessage{}, errors.New("ErrorMessage Not Found"))

			res, err := u.GetErrorMessageByCodeAndAppName(ctx, "400", "")

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r.Unset()
		})
	})
}
