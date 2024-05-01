package service

import (
	"anik/internal/model"
	"context"
)

type CompaniesService interface {
	Create(ctx context.Context, company *model.Company) (int, error)
	Get(ctx context.Context, id string) (*model.Company, error)
	GetAll(ctx context.Context) ([]model.Company, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, company *model.Company) error
}

type AnnouncementService interface {
	Create(ctx context.Context, announcement *model.Announcement) (int, error)
	Get(ctx context.Context, id string) (*model.Announcement, error)
	GetAll(ctx context.Context) ([]model.Announcement, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, announcement *model.Announcement) error
}

type CategoriesService interface {
	AddBusinessCategory(ctx context.Context, category *model.Category) (int, error)
	AddOfferCategory(ctx context.Context, category *model.Category) (int, error)
	GetBusinessCategories(ctx context.Context) ([]model.Category, error)
	GetOfficerCategories(ctx context.Context) ([]model.Category, error)
}
