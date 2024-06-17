package model

import "time"

type Company struct {
	ID          string     `db:"company_id" json:"company_id,omitempty" example:"0e3df004-ca0c-45a3-aeee-fa21c4aa3e4d"`
	Name        string     `db:"name" json:"name" example:"Company"`
	Description string     `db:"description" json:"description" example:"Company Description"`
	Address     string     `db:"address" json:"address" example:"Company Address"`
	LogoUrl     *string    `db:"logo_url" json:"logo_url,omitempty"`
	Latitude    float64    `db:"latitude" json:"latitude" example:"37.8483"`
	Longitude   float64    `db:"longitude" json:"longitude" example:"46.8483"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at,omitempty"`
	Categories  []string   `json:"category,omitempty" example:"Company Categories"`
	DistToUser  float64    `json:"distance_to_user,omitempty"`
}

func NewCompany() *Company {
	return &Company{}
}
