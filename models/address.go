package models

type Address struct {
	Id           uint                   `json:"id,omitempty"`
	PhoneNumber  string                 `json:"phone_number,omitempty"`
	AddressLine1 string                 `json:"address_line_1,omitempty"`
	AddressLine2 string                 `json:"address_line_2,omitempty"`
	UserId       uint                    `json:"user_id,omitempty"`
	GoogleUserId uint                    `json:"google_user_id,omitempty"`
}
