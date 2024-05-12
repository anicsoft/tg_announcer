package model

import "time"

type User struct {
	Id           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Username     string    `json:"username"`
	LanguageCode string    `json:"language_code"`
	Latitude     float32   `json:"latitude"`
	Longitude    float32   `json:"longitude"`
	UserType     int       `json:"user_type"`
	CreatedAt    time.Time `json:"created_at"`
}

func NewUser() *User {
	return &User{}
}
