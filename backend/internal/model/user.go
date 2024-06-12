package model

import "time"

const defaultUserType = "user"

type User struct {
	Id           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Username     string    `json:"username"`
	LanguageCode string    `json:"language_code"`
	Latitude     float32   `json:"latitude"`
	Longitude    float32   `json:"longitude"`
	UserType     string    `json:"user_type"`
	CompanyId    *string   `json:"company_id"`
	CreatedAt    time.Time `json:"created_at"`
}

func NewUser() *User {
	return &User{
		UserType: defaultUserType,
	}
}
