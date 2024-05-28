package model

import (
	"anik/internal/model"
	"time"
)

type AddAnnouncementResponse struct {
	ID int `json:"id"`
}

type AnnouncementsResponse struct {
	Announcements []model.Announcement `json:"announcements"`
}

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
	Categories []string `json:"categories,omitempty" example:"Special Offer"`
	StartDate  string   `json:"start_date,omitempty"`
	EndDate    string   `json:"end_date,omitempty"`
	PromoCode  bool     `json:"promo_code,omitempty" example:"true"`
	Latitude   float64  `json:"latitude,omitempty" example:"58.3854"`
	Longitude  float64  `json:"longitude,omitempty" example:"24.4971"`
}
