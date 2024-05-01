package api

import (
	"anik/internal/model"
	"context"
	"net/http"
)

func (a *BaseApi) AddOfferCategory(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category := model.NewCategory()
		err := a.Decode(r, &category)
		if err != nil {
			a.Error(w, http.StatusBadRequest, err)
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

func (a *BaseApi) OfferCategories(ctx context.Context) http.HandlerFunc {
	type response struct {
	}
	return func(w http.ResponseWriter, r *http.Request) {
		categories, err := a.categoriesService.GetOfficerCategories(ctx)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		a.Respond(w, http.StatusOK, categories)
	}
}

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
