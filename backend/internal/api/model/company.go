package model

type AddCompanyRequest struct {
	Name        string   `json:"name" example:"Tartu Bakery"`
	Description string   `json:"description" example:"Traditional Estonian bakery"`
	Address     string   `json:"address" example:"Tartu, Estonia"`
	Latitude    float64  `json:"latitude" example:"58.3780"`
	Longitude   float64  `json:"longitude" example:"26.7296"`
	Category    []string `json:"category" example:"Food & Drinks"`
}

type AddCompanyResponse struct {
	ID string `json:"id" example:"1"`
}
