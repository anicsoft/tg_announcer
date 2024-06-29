package model

import "time"

const defaultUserType = "user"

type User struct {
	ID           int       `json:"id" example:"12443543"`
	FirstName    string    `json:"first_name" example:"John"`
	LastName     string    `json:"last_name" example:"Doe"`
	Username     string    `json:"username" example:"johndoe"`
	LanguageCode string    `json:"language_code" example:"en"`
	UserType     string    `json:"user_type" example:"user|business"`
	CompanyId    *string   `json:"company_id" example:"0e3df004-ca0c-45a3-aeee-fa21c4aa3e4d"`
	CreatedAt    time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
}

func NewUser() *User {
	return &User{
		UserType: defaultUserType,
	}
}
