package request

type CreateRoleMenuRequest struct {
	RoleID    uint   `json:"role_id" validate:"required"`
	Menu      []uint `json:"menus" validate:"required"`
	CreatedBy uint   `json:"-"`
}
