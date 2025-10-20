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

func TestCreateCity(t *testing.T) {
	ctx := context.Background()

	Convey("CreateCity Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("Transaction", mock.Anything).Return(nil)
			r2 := m.On("CreateCity", mock.Anything).Return(model.City{}, nil)
			r3 := m.On("UpdateCity", mock.Anything).Return(model.City{}, nil)
			r4 := m.On("UpdateProvinceAddTotalCity", mock.Anything).Return(nil)

			res, err := u.CreateCity(ctx, request.CreateCityRequest{ProvinceID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
			r4.Unset()
		})
		Convey("fail CreateCity", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("CreateCity", mock.Anything).Return(model.City{}, errors.New("error"))
			r3 := m.On("UpdateCity", mock.Anything).Return(model.City{}, nil)

			res, err := u.CreateCity(ctx, request.CreateCityRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
		Convey("fail UpdateCity", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("CreateCity", mock.Anything).Return(model.City{}, nil)
			r3 := m.On("UpdateCity", mock.Anything).Return(model.City{}, errors.New("error"))

			res, err := u.CreateCity(ctx, request.CreateCityRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
	})
}

func TestGetCitys(t *testing.T) {
	ctx := context.Background()
	query := request.GetCityQuery{}

	Convey("GetCitys Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetCitys", query).Return([]model.City{{ID: 1, ProvinceID: 1, Name: "Kuta"}, {ID: 2, ProvinceID: 1, Name: "Lor"}}, nil)
			r2 := m.On("GetCityTotal", query).Return(uint(2), nil)

			res, err := u.GetCitys(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(res.Total, ShouldEqual, 2)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetCitys", func() {
			r1 := m.On("GetCitys", query).Return([]model.City{}, errors.New("Error GetCitys"))
			r2 := m.On("GetCityTotal", query).Return(uint(0), nil)

			res, err := u.GetCitys(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetCityTotal", func() {
			r1 := m.On("GetCitys", query).Return([]model.City{}, nil)
			r2 := m.On("GetCityTotal", query).Return(uint(0), errors.New("Error GetCityTotal"))

			res, err := u.GetCitys(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestUpdateCity(t *testing.T) {
	ctx := context.Background()

	Convey("UpdateCity Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetCityByID", mock.Anything).Return(model.City{ID: 1, Name: "Bali"}, nil)
			r2 := m.On("UpdateCity", mock.Anything).Return(model.City{ID: 1, Name: "Lor"}, nil)

			res, err := u.UpdateCity(ctx, request.UpdateCityRequest{ID: 1, Name: "Lor"})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)
			So(res.Name, ShouldEqual, "Lor")

			r1.Unset()
			r2.Unset()
		})
		Convey("fail UpdateCity", func() {
			r1 := m.On("GetCityByID", mock.Anything).Return(model.City{ID: 1}, nil)
			r2 := m.On("UpdateCity", mock.Anything).Return(model.City{}, errors.New("Error UpdateCity"))

			res, err := u.UpdateCity(ctx, request.UpdateCityRequest{ID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetCityByID", func() {
			r1 := m.On("GetCityByID", mock.Anything).Return(model.City{ID: 1}, errors.New("Error"))
			r2 := m.On("UpdateCity", mock.Anything).Return(model.City{}, nil)

			res, err := u.UpdateCity(ctx, request.UpdateCityRequest{ID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestGetCityByID(t *testing.T) {

	ctx := context.Background()

	Convey("GetCityDetail Usecase Scenario", t, func() {
		Convey("success", func() {
			r := m.On("GetCityByID", uint(1)).Return(model.City{ID: 1}, nil)

			res, err := u.GetCityDetail(ctx, uint(1))

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r.Unset()
		})
		Convey("fail", func() {
			r := m.On("GetCityByID", uint(0)).Return(model.City{ID: 0}, errors.New("City Not Found"))

			res, err := u.GetCityDetail(ctx, uint(0))

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r.Unset()
		})
	})
}

func TestDeleteCity(t *testing.T) {

	ctx := context.Background()

	Convey("DeleteCity Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("Transaction", mock.Anything).Return(nil)
			r2 := m.On("GetCityByID", uint(1)).Return(model.City{}, nil)
			r3 := m.On("DeleteCity", uint(1)).Return(nil)
			r4 := m.On("UpdateProvinceAddTotalCity", mock.Anything).Return(nil)

			err := u.DeleteCity(ctx, uint(1))

			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
			r4.Unset()
		})
		Convey("fail", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("GetCityByID", uint(0)).Return(model.City{}, errors.New("error"))
			r3 := m.On("DeleteCity", uint(0)).Return(nil)
			r4 := m.On("UpdateProvinceAddTotalCity", mock.Anything).Return(nil)

			err := u.DeleteCity(ctx, uint(0))

			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
			r4.Unset()
		})
	})
}
