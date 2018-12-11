package models

type User struct {
	Id             uint                   `json:"id,omitempty"`
	FirstName      string                 `json:"first_name,omitempty"`
	LastName       string                 `json:"last_name,omitempty"`
	Email          string                 `json:"email,omitempty"`
	Password       string                 `json:"password,omitempty"`
	SignedInSource string                 `json:"signed_in_source,omitempty"`  //  manual/google
	Address        []Address              `json:"addresses,omitempty"`
}

func (User) TableName() string {
	return "e_users"
}

