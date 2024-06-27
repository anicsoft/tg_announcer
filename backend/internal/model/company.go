package model

import "time"

type Company struct {
	ID           string                `db:"company_id" json:"company_id,omitempty" example:"0e3df004-ca0c-45a3-aeee-fa21c4aa3e4d"`
	Name         string                `db:"name" json:"name" example:"Company"`
	Description  string                `db:"description" json:"description" example:"Company Description"`
	Address      string                `db:"address" json:"address" example:"Company Address"`
	LogoUrl      *string               `db:"logo_url" json:"logo_url,omitempty"`
	Latitude     float64               `db:"latitude" json:"latitude" example:"37.8483"`
	Longitude    float64               `db:"longitude" json:"longitude" example:"46.8483"`
	UpdatedAt    *time.Time            `db:"updated_at" json:"updated_at,omitempty"`
	CreatedAt    *time.Time            `db:"created_at" json:"created_at,omitempty"`
	DeletedAt    *time.Time            `db:"deleted_at" json:"deleted_at,omitempty"`
	TelNumber    *string               `db:"tel_number" json:"tel_number,omitempty"`
	Email        *string               `db:"email" json:"email,omitempty"`
	Website      *string               `db:"website" json:"website,omitempty"`
	Facebook     *string               `db:"facebook" json:"facebook,omitempty"`
	Instagram    *string               `db:"instagram" json:"instagram,omitempty"`
	Telegram     *string               `db:"telegram" json:"telegram,omitempty"`
	Categories   []string              `json:"category,omitempty" example:"Company Categories"`
	DistToUser   float64               `json:"distance_to_user,omitempty"`
	WorkingHours []CompanyWorkingHours `json:"working_hours,omitempty"`
}

type CompanyWorkingHours struct {
	ID        string     `db:"id" json:"id,omitempty" example:"0e3df004-ca0c-45a3-aeee-fa21c4aa3e4d"`
	CompanyID string     `db:"company_id" json:"company_id,omitempty" example:"0e3df004-ca0c-45a3-aeee-fa21c4aa3e4d"`
	Day       string     `db:"day_of_week" json:"day_of_week" example:"Monday"`
	OpenTime  *time.Time `db:"open_time" json:"open_time,omitempty" example:"08:00:00"`
	CloseTime *time.Time `db:"close_time" json:"close_time,omitempty" example:"17:00:00"`
}

func NewCompany() *Company {
	return &Company{}
}
