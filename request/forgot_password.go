package request

type PhoneRequest struct {
	Phone string `json:"phone" validate:"required"`
}

type ChangePassword struct {
	Password             string `json:"password" validate:"required"`
	ConfirmationPassword string `json:"confirmation_password" validate:"required"`
	UserId               uint   `json:"-"`
}
