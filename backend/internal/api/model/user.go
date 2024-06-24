package model

type UpdateUserRequest struct {
	UserID    int    `json:"user_id" example:"0"`
	UserType  string `json:"user_type" example:"user|business"`
	CompanyId string `json:"company_id" example:"123442354"`
}
