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

func TestCreateMenu(t *testing.T) {
	ctx := context.Background()

	Convey("CreateMenu Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("Transaction", mock.Anything).Return(nil)
			r2 := m.On("CreateMenu", mock.Anything).Return(model.Menu{ParentMenuID: 0}, nil)
			r3 := m.On("UpdateMenu", mock.Anything).Return(model.Menu{}, nil)

			res, err := u.CreateMenu(ctx, request.CreateMenuRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
		Convey("all success", func() {
			r1 := m.On("Transaction", mock.Anything).Return(nil)
			r2 := m.On("CreateMenu", mock.Anything).Return(model.Menu{ParentMenuID: 1}, nil)
			r3 := m.On("GetLastOrderMenuByParentMenuID", mock.Anything).Return(model.Menu{}, nil)
			r4 := m.On("UpdateMenu", mock.Anything).Return(model.Menu{}, nil)

			res, err := u.CreateMenu(ctx, request.CreateMenuRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
			r4.Unset()
		})
		Convey("fail CreateMenu", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("CreateMenu", mock.Anything).Return(model.Menu{ParentMenuID: 0}, errors.New("error"))

			res, err := u.CreateMenu(ctx, request.CreateMenuRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetLastOrderMenuByParentMenuID", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("CreateMenu", mock.Anything).Return(model.Menu{ParentMenuID: 1}, nil)
			r3 := m.On("GetLastOrderMenuByParentMenuID", mock.Anything).Return(model.Menu{}, errors.New("error"))

			res, err := u.CreateMenu(ctx, request.CreateMenuRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
		Convey("fail UpdateMenu", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("error"))
			r2 := m.On("CreateMenu", mock.Anything).Return(model.Menu{ParentMenuID: 0}, nil)
			r3 := m.On("UpdateMenu", mock.Anything).Return(model.Menu{}, errors.New("error"))

			res, err := u.CreateMenu(ctx, request.CreateMenuRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
	})
}

func TestGetMenus(t *testing.T) {
	ctx := context.Background()
	query := request.GetMenuQuery{}

	Convey("GetMenus Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetMenus", query).Return([]model.Menu{}, nil)
			r2 := m.On("GetMenuTotal", query).Return(uint(2), nil)

			res, err := u.GetMenus(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetMenus", func() {
			r1 := m.On("GetMenus", query).Return([]model.Menu{}, errors.New("Error GetMenus"))
			r2 := m.On("GetMenuTotal", query).Return(uint(0), nil)

			res, err := u.GetMenus(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetMenuTotal", func() {
			r1 := m.On("GetMenus", query).Return([]model.Menu{}, nil)
			r2 := m.On("GetMenuTotal", query).Return(uint(0), errors.New("Error GetMenuTotal"))

			res, err := u.GetMenus(ctx, query)
			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestUpdateMenu(t *testing.T) {
	ctx := context.Background()

	Convey("UpdateMenu Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetMenuByID", mock.Anything).Return(model.Menu{}, nil)
			r2 := m.On("UpdateMenu", mock.Anything).Return(model.Menu{}, nil)

			res, err := u.UpdateMenu(ctx, request.UpdateMenuRequest{})

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetMenuByID", func() {
			r1 := m.On("GetMenuByID", mock.Anything).Return(model.Menu{ID: 1}, errors.New("Error"))
			r2 := m.On("UpdateMenu", mock.Anything).Return(model.Menu{}, nil)

			res, err := u.UpdateMenu(ctx, request.UpdateMenuRequest{ID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail UpdateMenu", func() {
			r1 := m.On("GetMenuByID", mock.Anything).Return(model.Menu{ID: 1}, nil)
			r2 := m.On("UpdateMenu", mock.Anything).Return(model.Menu{}, errors.New("Error"))

			res, err := u.UpdateMenu(ctx, request.UpdateMenuRequest{ID: 1})

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestUpdateOrderMenu(t *testing.T) {
	ctx := context.Background()

	Convey("UpdateOrderMenu Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("Transaction", mock.Anything).Return(nil)

			err := u.UpdateOrderMenu(ctx, request.UpdateOrderMenuRequest{})

			So(err, ShouldBeNil)

			r1.Unset()
		})
		Convey("fail", func() {
			r1 := m.On("Transaction", mock.Anything).Return(errors.New("ERROR"))

			err := u.UpdateOrderMenu(ctx, request.UpdateOrderMenuRequest{})

			So(err, ShouldNotBeNil)

			r1.Unset()
		})
	})
}

func TestGetMenuByID(t *testing.T) {

	ctx := context.Background()

	Convey("GetMenuDetail Usecase Scenario", t, func() {
		Convey("success", func() {
			r := m.On("GetMenuByID", uint(1)).Return(model.Menu{ID: 1}, nil)

			res, err := u.GetMenuDetail(ctx, uint(1))

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r.Unset()
		})
		Convey("fail", func() {
			r := m.On("GetMenuByID", uint(0)).Return(model.Menu{ID: 0}, errors.New("Menu Not Found"))

			res, err := u.GetMenuDetail(ctx, uint(0))

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r.Unset()
		})
	})
}

func TestGetMenuDetailWithChild(t *testing.T) {

	ctx := context.Background()

	Convey("GetMenuDetailWithChild Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetMenuByID", uint(1)).Return(model.Menu{ID: 1}, nil)
			r2 := m.On("GetMenusByParentID", uint(1)).Return([]model.Menu{}, nil)

			res, err := u.GetMenuDetailWithChild(ctx, uint(1))

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetMenuByID", func() {
			r1 := m.On("GetMenuByID", uint(1)).Return(model.Menu{}, errors.New("ERROR"))
			r2 := m.On("GetMenusByParentID", uint(1)).Return([]model.Menu{}, nil)

			res, err := u.GetMenuDetailWithChild(ctx, uint(1))

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
		Convey("fail GetMenusByParentID", func() {
			r1 := m.On("GetMenuByID", uint(1)).Return(model.Menu{}, nil)
			r2 := m.On("GetMenusByParentID", uint(1)).Return([]model.Menu{}, errors.New("ERROR"))

			res, err := u.GetMenuDetailWithChild(ctx, uint(1))

			So(res, ShouldNotBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
		})
	})
}

func TestGetMenusWithChild(t *testing.T) {

	ctx := context.Background()

	Convey("GetMenusWithChild Usecase Scenario", t, func() {
		Convey("success", func() {
			r1 := m.On("GetAllParentMenus", mock.Anything).Return([]model.Menu{{ID: 1}}, nil)
			r2 := m.On("GetMenuByID", uint(1)).Return(model.Menu{}, nil)
			r3 := m.On("GetRoleMenuByRoleIDAndMenuID", uint(1)).Return(model.RoleMenu{}, nil)
			r4 := m.On("GetMenusByParentID", uint(1)).Return([]model.Menu{}, nil)

			res, err := u.GetMenusWithChild(ctx, uint(1))

			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
			r4.Unset()
		})
		Convey("success 2", func() {
			r1 := m.On("GetAllParentMenus", mock.Anything).Return([]model.Menu{{ID: 1}, {ID: 2}}, nil)
			r2 := m.On("GetMenuByID", uint(1)).Return(model.Menu{}, errors.New("ERROR"))
			r3 := m.On("GetMenusByParentID", uint(1)).Return([]model.Menu{}, nil)
			r4 := m.On("GetMenuByID", uint(2)).Return(model.Menu{}, nil)
			r5 := m.On("GetRoleMenuByRoleIDAndMenuID", uint(0)).Return(model.RoleMenu{}, nil)
			r6 := m.On("GetMenusByParentID", uint(2)).Return([]model.Menu{}, nil)

			res, err := u.GetMenusWithChild(ctx, uint(0))

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
			r1 := m.On("GetAllParentMenus", mock.Anything).Return([]model.Menu{}, errors.New("Menu Not Found"))
			r2 := m.On("GetMenuByID", uint(1)).Return(model.Menu{ID: 1}, nil)
			r3 := m.On("GetMenusByParentID", uint(1)).Return([]model.Menu{}, nil)

			res, err := u.GetMenusWithChild(ctx, uint(0))

			So(res, ShouldBeNil)
			So(err, ShouldNotBeNil)

			r1.Unset()
			r2.Unset()
			r3.Unset()
		})
	})
}

func TestDeleteMenu(t *testing.T) {

	ctx := context.Background()

	Convey("DeleteMenu Usecase Scenario", t, func() {
		Convey("success", func() {
			r := m.On("DeleteMenu", uint(1)).Return(nil)

			err := u.DeleteMenu(ctx, uint(1))

			So(err, ShouldBeNil)

			r.Unset()
		})
		Convey("fail", func() {
			r := m.On("DeleteMenu", uint(0)).Return(errors.New("Menu Not Found"))

			err := u.DeleteMenu(ctx, uint(0))

			So(err, ShouldNotBeNil)

			r.Unset()
		})
	})
}
