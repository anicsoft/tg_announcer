package model

import (
	"time"
)

type Company struct {
	Id          string    `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Address     string    `db:"address" json:"address"`
	Latitude    float64   `db:"latitude" json:"latitude"`
	Longitude   float64   `db:"longitude" json:"longitude"`
	Who         string    `db:"who" json:"who"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

type CompanyWithDist struct {
	Company    Company `json:"company"`
	DistToUser float64 `json:"distance"`
}

func NewCompanyWithDistance(company *Company, distance float64) CompanyWithDist {
	return CompanyWithDist{
		Company:    *company,
		DistToUser: distance,
	}
}
