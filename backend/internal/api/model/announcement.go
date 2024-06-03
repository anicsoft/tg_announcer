package model

import (
	"anik/internal/model"
	"time"
)

type AddAnnouncementResponse struct {
	ID int `json:"id"`
}

type AnnouncementResponse struct {
	Announcements []model.Announcement `json:"announcements"`
}

type AddAnnouncement struct {
	CompanyID     int       `json:"company_id" example:"1"`
	Title         string    `json:"title" example:"We have free food!"`
	Categories    []string  `json:"categories" example:"Special Offer"`
	StartDateTime time.Time `json:"start_date_time" example:"2024-05-01 12:00:00.000000 +00:00"`
	EndDateTime   time.Time `json:"end_date_time" example:"2024-05-06 20:00:00.000000 +00:00"`
	PromoCode     *string   `json:"promo_code" example:"null"`
}

type Filter struct {
	Categories []string `json:"categories,omitempty" example:"Special Offer"`
	StartDate  string   `json:"start_date_time,omitempty" example:"2024-05-01 12:00:00.000000 +00:00"`
	EndDate    string   `json:"end_date_time,omitempty" example:"2024-05-06 20:00:00.000000 +00:00"`
	CreatedAt  string   `json:"created_at,omitempty" example:"2024-05-01 12:00:00.000000 +00:00"`
	PromoCode  bool     `json:"promo_code,omitempty" example:"true"`
	Latitude   float64  `json:"latitude,omitempty" example:"58.3854"`
	Longitude  float64  `json:"longitude,omitempty" example:"24.4971"`
	SortBy     string   `json:"sort_by,omitempty" example:"start_date_time"`
	SortOrder  string   `json:"sort_order,omitempty" example:"desc"`
	PageSize   int      `json:"page_size,omitempty"`
	Offset     int      `json:"offset,omitempty"`
}
