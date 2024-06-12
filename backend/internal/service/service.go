package service

import (
	apiModel "anik/internal/api/model"
	"anik/internal/model"
	"context"
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
	Get(ctx context.Context, parentId int) ([]string, error)
	Upload(ctx context.Context, parentId int, paths []string) error
	/*Delete(id string) error
	DeleteAll(parentId string) error*/
}
