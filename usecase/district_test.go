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

func TestCreateDistrict(t *testing.T) {
	ctx := context.Background()

	Convey("CreateDistrict Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("Transaction", mock.Anything).Return(nil)
			r2 := m.On("CreateDistrict", mock.Anything).Return(model.District{}, nil)
			r3 := m.On("UpdateDistrict", mock.Anything).Return(model.District{}, nil)
			r4 := m.On("GetCityByID", mock.Anything).Return(model.City{}, nil)
			r5 := m.On("UpdateProvinceAddTotalDistrict", mock.Anything).Return(nil)

			res, err := u.CreateDistrict(ctx, request.CreateDistrictRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
			r4.Unset()
			r5.Unset()
		})
		Convey("fail CreateDistrict", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("CreateDistrict", mock.Anything).Return(model.District{}, errors.New("error"))
			r3 := m.On("UpdateDistrict", mock.Anything).Return(model.District{}, nil)

			res, err := u.CreateDistrict(ctx, request.CreateDistrictRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
		Convey("fail UpdateDistrict", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("CreateDistrict", mock.Anything).Return(model.District{}, nil)
			r3 := m.On("UpdateDistrict", mock.Anything).Return(model.District{}, errors.New("error"))

			res, err := u.CreateDistrict(ctx, request.CreateDistrictRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
	})
}

func TestGetDistricts(t *testing.T) {
	ctx := context.Background()
	query := request.GetDistrictQuery{}

	Convey("GetDistricts Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetDistricts", query).Return([]model.District{}, nil)
			r2 := m.On("GetDistrictTotal", query).Return(uint(1), nil)

			res, err := u.GetDistricts(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetDistricts", func() {
			r1 := m.On("GetDistricts", query).Return([]model.District{}, errors.New("Error GetDistricts"))
			r2 := m.On("GetDistrictTotal", query).Return(uint(0), nil)

			res, err := u.GetDistricts(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetDistrictTotal", func() {
			r1 := m.On("GetDistricts", query).Return([]model.District{}, nil)
			r2 := m.On("GetDistrictTotal", query).Return(uint(0), errors.New("Error GetDistrictTotal"))

			res, err := u.GetDistricts(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestUpdateDistrict(t *testing.T) {
	ctx := context.Background()

	Convey("UpdateDistrict Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetDistrictByID", mock.Anything).Return(model.District{ID: 1}, nil)
			r2 := m.On("UpdateDistrict", mock.Anything).Return(model.District{ID: 1}, nil)

			res, err := u.UpdateDistrict(ctx, request.UpdateDistrictRequest{ID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail UpdateDistrict", func() {
			r1 := m.On("GetDistrictByID", mock.Anything).Return(model.District{ID: 1}, nil)
			r2 := m.On("UpdateDistrict", mock.Anything).Return(model.District{}, errors.New("Error UpdateDistrict"))

			res, err := u.UpdateDistrict(ctx, request.UpdateDistrictRequest{ID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetDistrictByID", func() {
			r1 := m.On("GetDistrictByID", mock.Anything).Return(model.District{ID: 1}, errors.New("error"))
			r2 := m.On("UpdateDistrict", mock.Anything).Return(model.District{}, errors.New("Error UpdateDistrict"))

			res, err := u.UpdateDistrict(ctx, request.UpdateDistrictRequest{ID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestGetDistrictByID(t *testing.T) {

	ctx := context.Background()

	Convey("GetDistrictDetail Usecase Scenario", t, func() {
		Convey("success", func() {
			r := m.On("GetDistrictByID", uint(1)).Return(model.District{ID: 1}, nil)

			res, err := u.GetDistrictDetail(ctx, uint(1))

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r.Unset()
		})
		Convey("fail", func() {
			r := m.On("GetDistrictByID", uint(0)).Return(model.District{ID: 0}, errors.New("District Not Found"))

			res, err := u.GetDistrictDetail(ctx, uint(0))

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r.Unset()
		})
	})
}

func TestDeleteDistrict(t *testing.T) {

	ctx := context.Background()

	Convey("DeleteDistrict Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("Transaction", mock.Anything).Return(nil)
			r2 := m.On("GetDistrictByID", uint(0)).Return(model.District{}, nil)
			r3 := m.On("DeleteDistrict", uint(0)).Return(nil)
			r4 := m.On("GetCityByID", uint(0)).Return(model.City{}, nil)
			r5 := m.On("UpdateProvinceAddTotalDistrict", uint(0)).Return(nil)

			err := u.DeleteDistrict(ctx, uint(0))

			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
			r4.Unset()
			r5.Unset()
		})
		Convey("fail", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("GetDistrictByID", uint(0)).Return(model.District{}, errors.New("error"))
			r3 := m.On("DeleteDistrict", uint(0)).Return(nil)

			err := u.DeleteDistrict(ctx, uint(0))

			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
	})
}
