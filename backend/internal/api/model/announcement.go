package model

import (
	"tg_announcer/internal/model"
	"time"
)

type AddAnnouncementResponse struct {
	ID string `json:"id" example:"0e3df004-ca0c-45a3-aeee-fa21c4aa3e4d"`
}

type AnnouncementResponse struct {
	Announcements []model.Announcement `json:"announcements"`
}

type AddAnnouncement struct {
	CompanyID     string    `json:"company_id" example:"0e3df004-ca0c-45a3-aeee-fa21c4aa3e4d"`
	Title         string    `json:"title" example:"We have free food!"`
	Content       string    `json:"content" example:"<h1>Hello, World!</h1>"`
	Categories    []string  `json:"categories" example:"Special Offer"`
	StartDateTime time.Time `json:"start_date_time" example:"2024-05-06T20:00:00.000000+00:00"`
	EndDateTime   time.Time `json:"end_date_time" example:"2024-05-06T20:00:00.000000+00:00"`
	PromoCode     *string   `json:"promo_code" example:"null"`
}

type Filter struct {
	CompanyID  string   `form:"company_id,omitempty" example:"0e3df004-ca0c-45a3-aeee-fa21c4aa3e4d"`
	Categories []string `form:"categories,omitempty" example:"Special Offer"`
	StartDate  string   `form:"start_date_time,omitempty" example:"2024-05-06T20:00:00.000000+00:00"`
	EndDate    string   `form:"end_date_time,omitempty" example:"2024-05-06T20:00:00.000000+00:00"`
	CreatedAt  string   `form:"created_at,omitempty" example:"2024-05-06T20:00:00.000000+00:00"`
	PromoCode  bool     `form:"promo_code,omitempty" example:"true"`
	Latitude   float64  `form:"latitude,omitempty" example:"58.3854"`
	Longitude  float64  `form:"longitude,omitempty" example:"24.4971"`
	SortBy     string   `form:"sort_by,omitempty" example:"start_date_time"`
	SortOrder  string   `form:"sort_order,omitempty" example:"desc"`
	PageSize   int      `json:"page_size,omitempty"`
	Offset     int      `form:"offset,omitempty"`
}
