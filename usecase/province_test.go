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

func TestCreateProvince(t *testing.T) {
	ctx := context.Background()

	Convey("CreateProvince Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("Transaction", mock.Anything).Return(nil)
			r2 := m.On("CreateProvince", mock.Anything).Return(model.Province{}, nil)
			r3 := m.On("UpdateProvince", mock.Anything).Return(model.Province{}, nil)

			res, err := u.CreateProvince(ctx, request.CreateProvinceRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
		Convey("fail CreateProvince", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("CreateProvince", mock.Anything).Return(model.Province{}, errors.New("error"))
			r3 := m.On("UpdateProvince", mock.Anything).Return(model.Province{}, nil)

			res, err := u.CreateProvince(ctx, request.CreateProvinceRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
		Convey("fail UpdateProvince", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("CreateProvince", mock.Anything).Return(model.Province{}, nil)
			r3 := m.On("UpdateProvince", mock.Anything).Return(model.Province{}, errors.New("error"))

			res, err := u.CreateProvince(ctx, request.CreateProvinceRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
	})
}

func TestGetProvinces(t *testing.T) {
	ctx := context.Background()
	query := request.GetProvinceQuery{}

	Convey("GetProvinces Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetProvinces", query).Return([]model.Province{}, nil)
			r2 := m.On("GetProvinceTotal", query).Return(uint(1), nil)

			res, err := u.GetProvinces(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetProvinces", func() {
			r1 := m.On("GetProvinces", query).Return([]model.Province{}, errors.New("Error GetProvinces"))
			r2 := m.On("GetProvinceTotal", query).Return(uint(0), errors.New("Error GetProvinceTotal"))

			res, err := u.GetProvinces(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetProvinceTotal", func() {
			r1 := m.On("GetProvinces", query).Return([]model.Province{}, nil)
			r2 := m.On("GetProvinceTotal", query).Return(uint(0), errors.New("Error GetProvinceTotal"))

			res, err := u.GetProvinces(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestUpdateProvince(t *testing.T) {
	ctx := context.Background()

	Convey("UpdateProvince Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetProvinceByID", mock.Anything).Return(model.Province{ID: 1}, nil)
			r2 := m.On("UpdateProvince", mock.Anything).Return(model.Province{ID: 1}, nil)

			res, err := u.UpdateProvince(ctx, request.UpdateProvinceRequest{ID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail UpdateProvince", func() {
			r1 := m.On("GetProvinceByID", mock.Anything).Return(model.Province{ID: 1}, nil)
			r2 := m.On("UpdateProvince", mock.Anything).Return(model.Province{}, errors.New("Error UpdateProvince"))

			res, err := u.UpdateProvince(ctx, request.UpdateProvinceRequest{ID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetProvinceByID", func() {
			r1 := m.On("GetProvinceByID", mock.Anything).Return(model.Province{ID: 1}, errors.New("error"))
			r2 := m.On("UpdateProvince", mock.Anything).Return(model.Province{}, errors.New("Error UpdateProvince"))

			res, err := u.UpdateProvince(ctx, request.UpdateProvinceRequest{ID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestGetProvinceByID(t *testing.T) {

	ctx := context.Background()

	Convey("GetProvinceDetail Usecase Scenario", t, func() {
		Convey("success", func() {
			r := m.On("GetProvinceByID", uint(1)).Return(model.Province{ID: 1}, nil)

			res, err := u.GetProvinceDetail(ctx, uint(1))

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r.Unset()
		})
		Convey("fail", func() {
			r := m.On("GetProvinceByID", uint(0)).Return(model.Province{ID: 0}, errors.New("Province Not Found"))

			res, err := u.GetProvinceDetail(ctx, uint(0))

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r.Unset()
		})
	})
}

func TestDeleteProvince(t *testing.T) {

	ctx := context.Background()

	Convey("DeleteProvince Usecase Scenario", t, func() {
		Convey("success", func() {
			r := m.On("DeleteProvince", uint(1)).Return(nil)

			err := u.DeleteProvince(ctx, uint(1))

			So(err, ShouldBeNil)

			r.Unset()
		})
		Convey("fail", func() {
			r := m.On("DeleteProvince", uint(0)).Return(errors.New("Province Not Found"))

			err := u.DeleteProvince(ctx, uint(0))

			So(err, ShouldNotBeNil)

			r.Unset()
		})
	})
}

func TestGetProvinceDetailWithTree(t *testing.T) {
	ctx := context.Background()

	Convey("GetProvinceDetailWithTree Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetProvinceByID", uint(1)).Return(model.Province{ID: 1}, nil)
			r2 := m.On("GetCitysByProvinceID", uint(1)).Return([]model.City{{ID: 1}, {ID: 2}}, nil)
			r3 := m.On("GetDistrictsByCityID", uint(1)).Return([]model.District{{ID: 1}}, nil)
			r4 := m.On("GetDistrictsByCityID", uint(2)).Return([]model.District{{ID: 2}}, nil)
			r5 := m.On("GetVillagesByDistrictID", uint(1)).Return([]model.Village{}, nil)
			r6 := m.On("GetVillagesByDistrictID", uint(2)).Return([]model.Village{}, nil)

			res, err := u.GetProvinceDetailWithTree(ctx, uint(1))

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
			r4.Unset()
			r5.Unset()
			r6.Unset()
		})
		Convey("fail", func() {
			r1 := m.On("GetProvinceByID", uint(1)).Return(model.Province{}, errors.New("error"))

			res, err := u.GetProvinceDetailWithTree(ctx, uint(1))

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
		})
	})
}
