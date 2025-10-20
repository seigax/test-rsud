package request

type RegisterFarmerRequest struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Phone           string `json:"phone" validate:"min=10"`
	Password        string `json:"password" validate:"min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=Password"`
	CreatedBy       uint   `json:"-"`
}

type LoginRequest struct {
	EmailOrPhone string `json:"email_or_phone" validate:"required"`
	Password     string `json:"password" validate:"min=6"`
}

type OTPRequest struct {
	EmailOrPhone string `json:"email_or_phone" validate:"required"`
	Otp          string `json:"otp" validate:"required"`
}
