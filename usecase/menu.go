package usecase

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/request"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/response"
)

func (usecase *Usecase) CreateMenu(ctx context.Context, payload request.CreateMenuRequest) (menu model.Menu, err error) {
	err = usecase.repo.Transaction(ctx, func(ctx context.Context) error {
		menu = model.Menu{
			ParentMenuID: payload.ParentMenuID,
			Name:         payload.Name,
			Level:        payload.Level,
			Url:          payload.Url,
			Icon:         payload.Icon,
			IsActiveFlag: payload.IsActiveFlag,
			CreatedAt:    time.Now(),
			CreatedBy:    payload.CreatedBy,
			UpdatedAt:    time.Now(),
			UpdatedBy:    payload.CreatedBy,
		}

		menu, err = usecase.repo.CreateMenu(ctx, menu)
		if err != nil {
			return err
		}

		if menu.IsChild() {
			lastMenu, err := usecase.repo.GetMenu(ctx, request.GetMenuQuery{
				Menu: model.Menu{ParentMenuID: menu.ParentMenuID},
				BasePaginateRequest: request.BasePaginateRequest{
					Sort: []string{"order_number-"},
				},
			})
			if err != nil {
				return err
			}

			menu.OrderNumber = lastMenu.OrderNumber + 1
		}

		menu.Code = fmt.Sprintf("MENU-%04d", menu.ID)

		_, err = usecase.repo.UpdateMenu(ctx, menu)
		if err != nil {
			return err
		}

		return nil
	})

	return
}

func (usecase *Usecase) GetMenus(ctx context.Context, query request.GetMenuQuery) (response.GetMenuResponse, error) {
	var res response.GetMenuResponse

	menus, err := usecase.repo.GetMenus(ctx, query)
	if err != nil {
		return res, err
	}

	total, err := usecase.repo.GetMenuTotal(ctx, query)
	if err != nil {
		return res, err
	}

	arrResp := []response.GetMenuWithParentNameResponse{}

	for _, v := range menus {
		respData := response.GetMenuWithParentNameResponse{
			Menu: v,
		}

		if v.IsChild() {
			parentMenu, _ := usecase.repo.GetMenu(ctx, request.GetMenuQuery{Menu: model.Menu{ID: v.ParentMenuID}})
			respData.ParentName = parentMenu.Name
		}

		arrResp = append(arrResp, respData)

	}

	res.Menus = arrResp
	res.Total = total
	res.Limit = query.Limit
	res.Offset = query.GetOffset()
	res.Page = query.Page

	return res, nil
}

func (usecase *Usecase) GetMenuDetail(ctx context.Context, ID uint) (model.Menu, error) {
	menu, err := usecase.repo.GetMenu(ctx, request.GetMenuQuery{Menu: model.Menu{ID: ID}})
	if err != nil {
		return menu, err
	}

	return menu, nil
}

func (usecase *Usecase) GetMenusWithChild(ctx context.Context, roleID uint) (resp []response.GetMenuDetailWithChildResponse, err error) {
	resp = []response.GetMenuDetailWithChildResponse{}
	allParent, err := usecase.repo.GetMenus(ctx, request.GetMenuQuery{Menu: model.Menu{Level: "Parent"}})
	if err != nil {
		return
	}

	for _, v := range allParent {

		childs, err := usecase.GetMenuDetailWithChildV2(ctx, v.ID, roleID)
		if err != nil {
			continue
		}

		resp = append(resp, childs)
	}

	return
}

func (usecase *Usecase) GetMenuDetailWithChildV2(ctx context.Context, ID uint, roleID uint) (resp response.GetMenuDetailWithChildResponse, err error) {
	parent, err := usecase.repo.GetMenu(ctx, request.GetMenuQuery{Menu: model.Menu{ID: ID}})
	if err != nil {
		return resp, err
	}

	resp.Menu = parent

	roleMenu, _ := usecase.repo.GetRoleMenuByRoleIDAndMenuID(ctx, roleID, ID)
	if roleMenu.ID == 0 {
		resp.IsSelected = "N"
	} else {
		resp.IsSelected = "Y"
	}

	childs, err := usecase.repo.GetMenus(ctx, request.GetMenuQuery{Menu: model.Menu{ParentMenuID: ID}})
	if err != nil {
		return resp, err
	}

	resp.Childs = []response.GetMenuDetailWithChildResponse{}

	for _, v := range childs {
		childResp, err := usecase.GetMenuDetailWithChildV2(ctx, v.ID, roleID)
		if err != nil {
			continue
		}

		resp.Childs = append(resp.Childs, childResp)
	}

	return resp, nil
}

func (usecase *Usecase) GetMenuDetailWithChild(ctx context.Context, ID uint) (resp response.GetMenuDetailWithChildResponse, err error) {
	parent, err := usecase.repo.GetMenu(ctx, request.GetMenuQuery{Menu: model.Menu{ID: ID}})
	if err != nil {
		return resp, err
	}

	resp.Menu = parent

	childs, err := usecase.repo.GetMenus(ctx, request.GetMenuQuery{Menu: model.Menu{ParentMenuID: ID}})
	if err != nil {
		return resp, err
	}

	resp.Childs = []response.GetMenuDetailWithChildResponse{}

	for _, v := range childs {
		childResp, err := usecase.GetMenuDetailWithChild(ctx, v.ID)
		if err != nil {
			continue
		}

		resp.Childs = append(resp.Childs, childResp)
	}

	return resp, nil
}

func (usecase *Usecase) UpdateMenu(ctx context.Context, payload request.UpdateMenuRequest) (model.Menu, error) {
	menu, err := usecase.repo.GetMenu(ctx, request.GetMenuQuery{Menu: model.Menu{ID: payload.ID}})
	if err != nil {
		return menu, err
	}

	menu.Name = payload.Name
	menu.Level = payload.Level
	menu.Url = payload.Url
	menu.Icon = payload.Icon
	menu.IsActiveFlag = payload.IsActiveFlag
	menu.UpdatedAt = time.Now()
	menu.UpdatedBy = payload.UpdatedBy

	res, err := usecase.repo.UpdateMenu(ctx, menu)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (usecase *Usecase) UpdateOrderMenu(ctx context.Context, payload request.UpdateOrderMenuRequest) error {
	err := usecase.repo.Transaction(ctx, func(ctx context.Context) error {
		for _, v := range payload.Menus {
			menu, err := usecase.repo.GetMenu(ctx, request.GetMenuQuery{Menu: model.Menu{ID: v.MenuID}})
			if err != nil {
				return err
			}

			menu.OrderNumber = v.OrderNumber
			_, err = usecase.repo.UpdateMenu(ctx, menu)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

func (usecase *Usecase) DeleteMenu(ctx context.Context, ID uint) error {
	err := usecase.repo.DeleteMenu(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
