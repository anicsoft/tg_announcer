package repository

import (
	"anik/internal/model"
	"context"
)

type CompaniesRepository interface {
	Create(ctx context.Context, company *model.Company) (int, error)
	Get(ctx context.Context, id string) (*model.Company, error)
	GetAll(ctx context.Context) ([]model.Company, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, company *model.Company) error
	AddCategory(ctx context.Context, category string, companyId int) error
	DeleteCategory(ctx context.Context, companyId int) error
}

type AnnouncementRepository interface {
	Create(ctx context.Context, announcement *model.Announcement) (int, error)
	Get(ctx context.Context, id string) (*model.Announcement, error)
	GetAll(ctx context.Context) ([]model.Announcement, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, company *model.Announcement) error
	AddCategory(ctx context.Context, category string, announcementId int) error
	//DeleteCategory(ctx context.Context, companyId int) error
}

type CategoriesRepository interface {
	AddOfferCategory(ctx context.Context, category string) (int, error)
	AddBusinessCategory(ctx context.Context, category string) (int, error)
	GetOfferCategories(ctx context.Context) ([]model.Category, error)
	GetBusinessCategories(ctx context.Context) ([]model.Category, error)
}
