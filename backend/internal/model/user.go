package model

import "time"

const defaultUserType = "user"

type User struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Username     string    `json:"username"`
	LanguageCode string    `json:"language_code"`
	UserType     string    `json:"user_type"`
	CompanyId    *string   `json:"company_id"`
	CreatedAt    time.Time `json:"created_at"`
}

func NewUser() *User {
	return &User{
		UserType: defaultUserType,
	}
}
