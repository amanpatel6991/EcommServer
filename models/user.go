package models

type User struct {
	Id             uint                   `json:"id,omitempty"`
	FirstName      string                 `json:"first_name,omitempty"`
	LastName       string                 `json:"last_name,omitempty"`
	Email          string                 `json:"email,omitempty"`
	Password       string                 `json:"password,omitempty"`
	SignedInSource string                 `json:"signed_in_source,omitempty"`
	PhoneNumber    string                 `json:"phone_number,omitempty"`
	AddressLine1   string                 `json:"address_line_1,omitempty"`
	AddressLine2   string                 `json:"address_line_2,omitempty"`
	//UserType     string                 `json:"usertype,omitempty"`
	//OrgId        uint                   `json:"orgid,omitempty"`
}

func (User) TableName() string {
	return "e_users"
}

