package response

import (
	"time"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/model"
)

type LoginResponse struct {
	ID               uint       `json:"id"`
	Name             string     `json:"name"`
	Email            string     `json:"email"`
	PhotoURL         string     `json:"photo_url"`
	ChangePasswordAt time.Time  `json:"change_password_at"`
	TncAcceptedAt    time.Time  `json:"tnc_accepted_at"`
	Role             model.Role `json:"role"`
	Token            string     `json:"token"`
}
