package model

import "tg_announcer/internal/model"

type UpdateUserRequest struct {
	UserID    int    `json:"user_id" example:"0"`
	UserType  string `json:"user_type" example:"user|business"`
	CompanyId string `json:"company_id" example:"123442354"`
}

type AddFavoriteRequest struct {
	CompanyID string `json:"company_id" example:"0e3df004-ca0c-45a3-aeee-fa21c4aa3e4d"`
}

type AddFavoriteResponse struct {
}

type FavoritesResponse struct {
	Companies []model.Company `json:"companies"`
}

type DeleteFavoriteRequest struct {
	CompanyID string `json:"company_id" example:"0e3df004-ca0c-45a3-aeee-fa21c4aa3e4d"`
}

type ListUserResponse struct {
	Users []model.User `json:"users"`
}
