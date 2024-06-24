package service

import (
	"context"
	apiModel "tg_announcer/internal/api/model"
	"tg_announcer/internal/model"
)

type CompaniesService interface {
	Create(ctx context.Context, company *model.Company) (string, error)
	GetByID(ctx context.Context, id string) (*model.Company, error)
	GetAll(ctx context.Context) ([]model.Company, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, company *model.Company) error
}

type AnnouncementService interface {
	Create(ctx context.Context, announcement *model.Announcement) (string, error)
	Get(ctx context.Context, id string) (*model.Announcement, error)
	GetAll(ctx context.Context, filters apiModel.Filter) ([]model.Announcement, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, announcement *model.Announcement) error
}

type CategoriesService interface {
	AddBusinessCategory(ctx context.Context, category *model.Category) (int, error)
	AddOfferCategory(ctx context.Context, category *model.Category) (int, error)
	GetBusinessCategories(ctx context.Context) ([]model.Category, error)
	GetOfficerCategories(ctx context.Context) ([]model.Category, error)
}

type UsersService interface {
	AddUser(ctx context.Context, user *model.User) (int, error)
	Exists(ctx context.Context, id int) (bool, error)
	GetByID(ctx context.Context, id int) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
}

type ImageService interface {
	GetLogo(ctx context.Context, companyId string) (string, error)
	UploadLogo(ctx context.Context, companyId string, path string) error
	GetAnnouncPictures(ctx context.Context, announcementId string) ([]string, error)
	UploadAnnouncPictures(ctx context.Context, announcementId string, paths []string) error
}
