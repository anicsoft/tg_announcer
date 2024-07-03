package model

import (
	"time"
)

type Announcement struct {
	AnnouncementID string    `db:"announcement_id" json:"announcement_id,omitempty"`
	CompanyID      string    `db:"company_id" json:"company_id,omitempty" example:"0e3df004-ca0c-45a3-aeee-fa21c4aa3e4d"`
	Title          string    `db:"title" json:"title" example:"We have free food!"`
	Content        string    `db:"content" json:"content" example:"<h1>Hello World!</h1>"`
	PromoCode      *string   `db:"promo_code" json:"promo_code" example:"PROMO|null"`
	PictureUrl     *string   `json:"picture_url,omitempty"`
	StartDateTime  time.Time `db:"start_date_time" json:"start_date_time" example:"2024-05-06T20:00:00.000000+00:00"`
	EndDateTime    time.Time `db:"end_date_time" json:"end_date_time" example:"2024-05-01T12:00:00.000000+00:00"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	Active         bool      `db:"active" json:"active" example:"true"`
	Categories     []string  `json:"categories" example:"Special Offer"`
	Company        Company   `db:"company" json:"company,omitempty"`
	Distance       float64   `db:"distance" json:"distance,omitempty" example:"99"`
}

func NewAnnouncement() *Announcement {
	return &Announcement{
		CreatedAt: time.Now(),
	}
}
