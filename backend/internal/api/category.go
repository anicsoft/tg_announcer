package api

import (
	"anik/internal/model"
	"context"
	"net/http"
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
func (a *BaseApi) AddOfferCategory(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category := model.NewCategory()
		err := a.Decode(r, &category)
		if err != nil {
			a.Error(w, http.StatusBadRequest, ErrDecodeBody)
			return
		}

		_, err = a.categoriesService.AddOfferCategory(ctx, category)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		a.Respond(w, http.StatusCreated, category)
	}
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
func (a *BaseApi) AddBusinessCategory(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category := model.NewCategory()
		err := a.Decode(r, &category)
		if err != nil {
			a.Error(w, http.StatusBadRequest, err)
			return
		}

		_, err = a.categoriesService.AddBusinessCategory(ctx, category)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		a.Respond(w, http.StatusCreated, category)
	}
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
func (a *BaseApi) OfferCategories(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categories, err := a.categoriesService.GetOfficerCategories(ctx)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		a.Respond(w, http.StatusOK, categories)
	}
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
func (a *BaseApi) BusinessCategories(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categories, err := a.categoriesService.GetBusinessCategories(ctx)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		a.Respond(w, http.StatusOK, categories)
	}
}
