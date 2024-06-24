package api

import (
	"errors"
	"tg_announcer/internal/model"

	"github.com/gin-gonic/gin"
)

// AddOfferCategory godoc
//
//	@Summary		Add offer category
//	@Description	Add new offer category
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"tma initData"
//	@Param			announcement	body		model.AddCategory	true	"request body"
//	@Success		201				{object}	model.Category
//	@Failure		400				{object}	HttpError	"failed to decode body"
//	@Failure		500				{object}	HttpError	"internal error"
//	@Router			/categories/offer [post]
func (a *BaseApi) AddOfferCategory(ctx *gin.Context) {
	category := model.NewCategory()
	err := ctx.ShouldBind(&category)
	if err != nil {
		StatusBadRequest(ctx, errors.Join(ErrDecodeBody, err))
		return
	}

	_, err = a.categoriesService.AddOfferCategory(ctx, category)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusCreated(ctx, category)
}

// AddBusinessCategory godoc
//
//	@Summary		Add business category
//	@Description	Add new business category
//	@Tags			categories
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"tma initData"
//	@Param			announcement	body		model.AddCategory	true	"request body"
//	@Success		201				{object}	model.Category
//	@Failure		400				{object}	HttpError	"failed to decode body"
//	@Failure		500				{object}	HttpError	"internal error"
//	@Router			/categories/business [post]
func (a *BaseApi) AddBusinessCategory(ctx *gin.Context) {
	category := model.NewCategory()
	err := ctx.ShouldBind(&category)
	if err != nil {
		StatusBadRequest(ctx, errors.Join(ErrDecodeBody, err))
		return
	}

	_, err = a.categoriesService.AddBusinessCategory(ctx, category)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusCreated(ctx, category)
}

// OfferCategories godoc
//
//	@Summary		List offer categories
//	@Description	List offer categories
//	@Tags			categories
//	@Produce		json
//	@Success		200	{object}	[]model.Category
//	@Failure		500	{object}	HttpError	"internal error"
//	@Router			/categories/offer [get]
func (a *BaseApi) OfferCategories(ctx *gin.Context) {
	categories, err := a.categoriesService.GetOfficerCategories(ctx)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusOK(ctx, categories)
}

// BusinessCategories godoc
//
//	@Summary		List business categories
//	@Description	List business categories
//	@Tags			categories
//	@Produce		json
//	@Success		200	{object}	[]model.Category
//	@Failure		500	{object}	HttpError	"internal error"
//	@Router			/categories/business [get]
func (a *BaseApi) BusinessCategories(ctx *gin.Context) {
	categories, err := a.categoriesService.GetBusinessCategories(ctx)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusOK(ctx, categories)
}
