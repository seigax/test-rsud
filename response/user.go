package response

import "gitlab.com/erloom.id/libraries/go/backend-skeleton/model"

type GetUserResponse struct {
	BasePaginateResponse
	Users []GetUserWithRolesResponse `json:"users"`
}

type GetUserWithRolesResponse struct {
	model.User
	Roles    string `json:"roles"`
	RoleType string `json:"role_type"`
}
