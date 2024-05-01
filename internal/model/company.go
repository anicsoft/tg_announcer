package model

type Company struct {
	Id          int      `db:"company_id" json:"company_id"`
	Name        string   `db:"name" json:"name"`
	Description string   `db:"description" json:"description"`
	Address     string   `db:"address" json:"address"`
	Category    []string `json:"category"`
}

func NewCompany() *Company {
	return &Company{}
}
