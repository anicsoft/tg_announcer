package api

import (
	"net/http"
	"tg_announcer/internal/service"

	"github.com/gin-gonic/gin"
)

type Api interface {
	Notify(ctx *gin.Context)
	Update(ctx *gin.Context)
	AddUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	AddCompany(ctx *gin.Context)
	GetCompanyByID(ctx *gin.Context)
	AddAnnouncement(ctx *gin.Context)
	GetAnnouncement(ctx *gin.Context)
	AddOfferCategory(ctx *gin.Context)
	AddBusinessCategory(ctx *gin.Context)
	OfferCategories(ctx *gin.Context)
	BusinessCategories(ctx *gin.Context)
	Announcements(ctx *gin.Context)
	UploadLogo(ctx *gin.Context)
	UploadImage(ctx *gin.Context)
}

type Response struct {
	Data interface{} `json:"data"`
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
		"message": err,
		"code":    "INTERNAL_SERVER_ERROR",
	})
}

func StatusBadRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": err,
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
