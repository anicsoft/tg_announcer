package model

import "time"

type Announcement struct {
	AnnouncementID int       `db:"announcement_id" json:"announcement_id"`
	CompanyID      int       `db:"company_id" json:"company_id" example:"some id"`
	Title          string    `db:"title" json:"title" example:"We have free food!"`
	Categories     []string  `json:"categories" example:"Special Offer"`
	StartDate      time.Time `db:"start_date" json:"start_date" example:"2024-07-20T00:00:00Z"`
	EndDate        time.Time `db:"end_date" json:"end_date" example:"2024-07-21T00:00:00Z"`
	StartTime      time.Time `db:"start_time" json:"start_time" example:"2000-01-01T21:00:00Z"`
	EndTime        time.Time `db:"end_time" json:"end_time" example:"2000-01-01T02:00:00Z"`
	PromoCode      *string   `db:"promo_code" json:"promo_code" example:"PROMO|null"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}

func NewAnnouncement() *Announcement {
	return &Announcement{
		CreatedAt: time.Now(),
	}
}
