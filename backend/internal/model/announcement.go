package model

import "time"

type Announcement struct {
	AnnouncementID int       `db:"announcement_id" json:"announcement_id"`
	CompanyID      int       `db:"company_id" json:"company_id"`
	Title          string    `db:"title" json:"title"`
	Categories     []string  `json:"categories"`
	StartDate      time.Time `db:"start_date" json:"start_date"`
	EndDate        time.Time `db:"end_date" json:"end_date"`
	StartTime      time.Time `db:"start_time" json:"start_time"`
	EndTime        time.Time `db:"end_time" json:"end_time"`
	PromoCode      *string   `db:"promo_code" json:"promo_code"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}

func NewAnnouncement() *Announcement {
	return &Announcement{
		CreatedAt: time.Now(),
	}
}
