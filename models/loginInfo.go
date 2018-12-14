package models

type LoginInfo struct {
	Email          string              `json:"email,omitempty"`
	Password       string                 `json:"password,omitempty"`
}

