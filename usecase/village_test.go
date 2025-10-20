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

func TestCreateVillage(t *testing.T) {
	ctx := context.Background()

	Convey("CreateVillage Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("Transaction", mock.Anything).Return(nil)
			r2 := m.On("CreateVillage", mock.Anything).Return(model.Village{}, nil)
			r3 := m.On("UpdateVillage", mock.Anything).Return(model.Village{}, nil)
			r4 := m.On("GetDistrict", mock.Anything).Return(model.District{}, nil)
			r5 := m.On("GetCity", mock.Anything).Return(model.City{}, nil)
			r6 := m.On("UpdateProvinceAddTotalVillage", mock.Anything).Return(nil)

			res, err := u.CreateVillage(ctx, request.CreateVillageRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
			r4.Unset()
			r5.Unset()
			r6.Unset()
		})
		Convey("fail CreateVillage", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("CreateVillage", mock.Anything).Return(model.Village{}, errors.New("error"))
			r3 := m.On("UpdateVillage", mock.Anything).Return(model.Village{}, nil)

			res, err := u.CreateVillage(ctx, request.CreateVillageRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
		Convey("fail UpdateVillage", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("CreateVillage", mock.Anything).Return(model.Village{}, nil)
			r3 := m.On("UpdateVillage", mock.Anything).Return(model.Village{}, errors.New("error"))

			res, err := u.CreateVillage(ctx, request.CreateVillageRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
	})
}

func TestGetVillages(t *testing.T) {
	ctx := context.Background()
	query := request.GetVillageQuery{}

	Convey("GetVillages Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetVillages", mock.Anything).Return([]model.Village{}, nil)
			r2 := m.On("GetVillageTotal", mock.Anything).Return(uint(1), nil)

			res, err := u.GetVillages(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetVillages", func() {
			r1 := m.On("GetVillages", mock.Anything).Return([]model.Village{}, errors.New("Error GetVillages"))
			r2 := m.On("GetVillageTotal", mock.Anything).Return(uint(0), errors.New("Error GetVillageTotal"))

			res, err := u.GetVillages(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetVillageTotal", func() {
			r1 := m.On("GetVillages", mock.Anything).Return([]model.Village{}, nil)
			r2 := m.On("GetVillageTotal", mock.Anything).Return(uint(0), errors.New("Error GetVillageTotal"))

			res, err := u.GetVillages(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestUpdateVillage(t *testing.T) {
	ctx := context.Background()

	Convey("UpdateVillage Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetVillageByID", mock.Anything).Return(model.Village{ID: 1}, nil)
			r2 := m.On("UpdateVillage", mock.Anything).Return(model.Village{ID: 1}, nil)

			res, err := u.UpdateVillage(ctx, request.UpdateVillageRequest{ID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail UpdateVillage", func() {
			r1 := m.On("GetVillageByID", mock.Anything).Return(model.Village{ID: 1}, nil)
			r2 := m.On("UpdateVillage", mock.Anything).Return(model.Village{}, errors.New("Error UpdateVillage"))

			res, err := u.UpdateVillage(ctx, request.UpdateVillageRequest{ID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetVillage", func() {
			r1 := m.On("GetVillageByID", mock.Anything).Return(model.Village{ID: 1}, errors.New("error"))
			r2 := m.On("UpdateVillage", mock.Anything).Return(model.Village{}, errors.New("Error UpdateVillage"))

			res, err := u.UpdateVillage(ctx, request.UpdateVillageRequest{ID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestGetVillage(t *testing.T) {

	ctx := context.Background()

	Convey("GetVillageDetail Usecase Scenario", t, func() {
		Convey("success", func() {
			r := m.On("GetVillageByID", uint(1)).Return(model.Village{ID: 1}, nil)

			res, err := u.GetVillageDetail(ctx, uint(1))

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r.Unset()
		})
		Convey("fail", func() {
			r := m.On("GetVillageByID", uint(0)).Return(model.Village{ID: 0}, errors.New("Village Not Found"))

			res, err := u.GetVillageDetail(ctx, uint(0))

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r.Unset()
		})
	})
}

func TestDeleteVillage(t *testing.T) {

	ctx := context.Background()

	Convey("DeleteVillage Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("Transaction", mock.Anything).Return(nil)
			r2 := m.On("GetVillageByID", uint(0)).Return(model.Village{}, nil)
			r3 := m.On("DeleteVillage", uint(0)).Return(nil)
			r4 := m.On("GetDistrict", uint(0)).Return(model.District{}, nil)
			r5 := m.On("GetCity", uint(0)).Return(model.City{}, nil)
			r6 := m.On("UpdateProvinceAddTotalVillage", uint(0)).Return(nil)

			err := u.DeleteVillage(ctx, uint(0))

			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
			r4.Unset()
			r5.Unset()
			r6.Unset()
		})
		Convey("fail", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("GetVillage", uint(0)).Return(model.Village{}, errors.New("error"))

			err := u.DeleteVillage(ctx, uint(0))

			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}
