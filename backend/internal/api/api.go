package api

import (
	"net/http"
	"tg_announcer/internal/service"

	"github.com/gin-gonic/gin"
)

type Api interface {
	AnnouncementApi
	CompanyApi
	UserApi
	CategoryApi
	Notify(ctx *gin.Context)
}

type AnnouncementApi interface {
	AddAnnouncement(ctx *gin.Context)
	GetAnnouncement(ctx *gin.Context)
	GetAnnouncements(ctx *gin.Context)
	UpdateAnnouncements(ctx *gin.Context)
	UploadImage(ctx *gin.Context)
}

type CompanyApi interface {
	AddCompany(ctx *gin.Context)
	GetCompanyByID(ctx *gin.Context)
	UpdateCompany(ctx *gin.Context)
	DeleteCompany(ctx *gin.Context)
	ListCompanies(ctx *gin.Context)
	UploadLogo(ctx *gin.Context)
}

type UserApi interface {
	GetUser(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type CategoryApi interface {
	AddOfferCategory(ctx *gin.Context)
	AddBusinessCategory(ctx *gin.Context)
	OfferCategories(ctx *gin.Context)
	BusinessCategories(ctx *gin.Context)
}

type BaseApi struct {
	companiesService    service.CompaniesService
	announcementService service.AnnouncementService
	categoriesService   service.CategoriesService
	userService         service.UsersService
	imageService        service.ImageService
}

func New(
	companiesService service.CompaniesService,
	announcementService service.AnnouncementService,
	categoriesService service.CategoriesService,
	userService service.UsersService,
	imageService service.ImageService,
) Api {
	return &BaseApi{
		companiesService:    companiesService,
		announcementService: announcementService,
		categoriesService:   categoriesService,
		userService:         userService,
		imageService:        imageService,
	}
}

func StatusInternalServerError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
		"code":    "INTERNAL_SERVER_ERROR",
	})
}

func StatusBadRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": err.Error(),
		"code":    "BAD_REQUEST",
	})
}

func StatusOK(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, data)
}

func StatusCreated(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusCreated, data)
}

func StatusAccepted(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, nil)
}
