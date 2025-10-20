package response

type UserPhoneResponse struct {
	ID           uint   `json:"id"`
	UserId       uint   `json:"user_id"`
	PhoneNumber  string `json:"phone_number"`
	IsActiveFlag string `json:"is_active_flag"`
}
