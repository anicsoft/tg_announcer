package model

type Company struct {
	Id          int      `db:"company_id" json:"company_id,omitempty" example:"1"`
	Name        string   `db:"name" json:"name" example:"Company"`
	Description string   `db:"description" json:"description" example:"Company Description"`
	Address     string   `db:"address" json:"address" example:"Company Address"`
	Latitude    float64  `db:"latitude" json:"latitude" example:"37.8483"`
	Longitude   float64  `db:"longitude" json:"longitude" example:"46.8483"`
	Category    []string `json:"category,omitempty" example:"Company Category"`
	DistToUser  float64  `json:"distance_to_user"`
}

func NewCompany() *Company {
	return &Company{}
}
