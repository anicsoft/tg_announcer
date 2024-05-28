package model

import "time"

type AddAnnouncement struct {
	CompanyID  int       `json:"company_id" example:"1"`
	Title      string    `json:"title" example:"We have free food!"`
	Categories []string  `json:"categories" example:"Special Offer"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	PromoCode  *string   `json:"promo_code" example:"null"`
}

type Filter struct {
	Category  string `json:"category,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	PromoCode string `json:"promo_code,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}
