package response

type UserFarmerResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	PhotoURL string `json:"photo_url"`
	Token    string `json:"token"`
}
