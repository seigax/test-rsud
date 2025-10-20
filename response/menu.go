package response

import "gitlab.com/erloom.id/libraries/go/backend-skeleton/model"

type GetMenuResponse struct {
	BasePaginateResponse
	Menus []GetMenuWithParentNameResponse `json:"menus"`
}

type GetMenuWithParentNameResponse struct {
	model.Menu
	ParentName string `json:"parent_name"`
}

type GetMenuDetailWithChildResponse struct {
	model.Menu
	IsSelected string                           `json:"is_selected,omitempty"`
	Childs     []GetMenuDetailWithChildResponse `json:"childs"`
}
