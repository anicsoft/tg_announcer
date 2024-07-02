package repository

import (
	"context"
	apiModel "tg_announcer/internal/api/model"
	"tg_announcer/internal/model"
)

type CompaniesRepository interface {
	Create(ctx context.Context, company *model.Company) (string, error)
	Get(ctx context.Context, id string) (*model.Company, error)
	GetAll(ctx context.Context) ([]model.Company, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, company *model.Company) error
	AddCategory(ctx context.Context, category string, companyId string) error
	DeleteCategory(ctx context.Context, companyId string) error
}

type AnnouncementRepository interface {
	Create(ctx context.Context, announcement *model.Announcement) (string, error)
	Get(ctx context.Context, id string) (*model.Announcement, error)
	GetAll(ctx context.Context, filter apiModel.Filter) ([]model.Announcement, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, company *model.Announcement) error
	AddCategory(ctx context.Context, category string, announcementId string) error
	//DeleteCategory(ctx context.Context, companyId int) error
}

type CategoriesRepository interface {
	AddOfferCategory(ctx context.Context, category string) (int, error)
	AddBusinessCategory(ctx context.Context, category string) (int, error)
	GetOfferCategories(ctx context.Context) ([]model.Category, error)
	GetBusinessCategories(ctx context.Context) ([]model.Category, error)
}

type UsersRepository interface {
	Create(ctx context.Context, user *model.User) (int, error)
	GetByID(ctx context.Context, id int) (*model.User, error)
	GetByUsername(ctx context.Context, id string) (*model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, user *model.User) error
	Exists(ctx context.Context, id int) (bool, error)
	AddFavoriteCompany(ctx context.Context, userId int, companyId string) error
	GetFavoriteCompanies(ctx context.Context, userId int) ([]string, error)
	DeleteFavoriteCompany(ctx context.Context, userId int, companyId string) error
	IsFavoriteCompany(ctx context.Context, userId int, companyId string) (bool, error)
}

type ImageRepository interface {
	AddLogo(ctx context.Context, companyId string, path string) (string, error)
	GetLogo(ctx context.Context, companyId string) (string, error)
	AddAnnouncementPictures(ctx context.Context, announcementId string, paths []string) (string, error)
	GetAnnouncementPictures(ctx context.Context, announcementId string) ([]string, error)
}
