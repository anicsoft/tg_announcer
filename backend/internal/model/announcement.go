package model

import (
	"time"
)

type Announcement struct {
	AnnouncementID int       `db:"announcement_id" json:"announcement_id,omitempty"`
	CompanyID      int       `db:"company_id" json:"company_id,omitempty" example:"1"`
	Title          string    `db:"title" json:"title" example:"We have free food!"`
	Content        string    `db:"content" json:"content" example:"<h1>Hello World!</h1>"`
	Categories     []string  `json:"categories" example:"Special Offer"`
	StartDateTime  time.Time `db:"start_date_time" json:"start_date_time" example:"2024-05-01 12:00:00.000000 +00:00"`
	EndDateTime    time.Time `db:"end_date_time" json:"end_date_time" example:"2024-06-01 12:00:00.000000 +00:00"`
	PromoCode      *string   `db:"promo_code" json:"promo_code" example:"PROMO|null"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	Company        Company   `db:"company" json:"company,omitempty"`
	Distance       float64   `db:"distance" json:"distance,omitempty" example:"99"`
}

func NewAnnouncement() *Announcement {
	return &Announcement{
		CreatedAt: time.Now(),
	}
}
