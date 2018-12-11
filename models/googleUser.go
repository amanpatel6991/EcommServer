package models

type GoogleUser struct {
	Id      uint                   `json:"id,omitempty"`
	Token   string                 `json:"token,omitempty"`
	Name    string                 `json:"name,omitempty"`
	Email   string                 `json:"email,omitempty"`
	Address []Address              `json:"addresses,omitempty"`
}

func (GoogleUser) TableName() string {
	return "e_google_user"
}