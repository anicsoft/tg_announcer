package api

import (
	apiModel "anik/internal/api/model"
	"anik/internal/model"
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// AddCompany godoc
//
//	@Summary		Adds company to database
//	@Description	Only for admins
//	@Tags			companies
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string					true	"tma initData"
//	@Param			announcement	body		model.AddCompanyRequest	true	"request body"
//	@Success		201				{object}	model.AddCompanyResponse
//	@Failure		400				{object}	HttpError	"failed to decode body"
//	@Failure		500				{object}	HttpError	"internal error"
//	@Router			/companies [post]
func (a *BaseApi) AddCompany(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		company := model.NewCompany()

		err := a.Decode(r, &company)
		if err != nil {
			a.Error(w, http.StatusBadRequest, errors.Join(ErrDecodeBody, err))
			return
		}

		id, err := a.companiesService.Create(ctx, company)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		a.Respond(w, http.StatusCreated, apiModel.AddCompanyResponse{ID: id})
	}
}

// GetCompanyByID godoc
//
//	@Summary		Returns company by ID
//	@Description	Only for admins
//	@Tags			companies
//	@Produce		json
//	@Param			id	path		int	true	"request body"
//	@Success		200				{object}	model.Company
//	@Failure		400				{object}	HttpError	"failed to decode body"
//	@Failure		500				{object}	HttpError	"internal error"
//	@Router			/companies/{id} [get]
func (a *BaseApi) GetCompanyByID(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			a.Error(w, http.StatusBadRequest, fmt.Errorf("empty id"))
			return
		}

		company, err := a.companiesService.GetByID(ctx, id)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		a.Respond(w, http.StatusOK, company)
	}
}

//func (a *Api) Delete(ctx context.Context) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		id := chi.URLParam(r, "id")
//		if id == "" {
//			a.Error(w, http.StatusBadRequest, fmt.Errorf("empty id"))
//			return
//		}
//
//		intId, err := strconv.Atoi(id)
//		if err != nil {
//			a.Error(w, http.StatusInternalServerError, err)
//			return
//		}
//
//		err = a.companiesService.Delete(ctx, intId)
//		if err != nil {
//			a.Error(w, http.StatusInternalServerError, err)
//			return
//		}
//
//		a.Respond(w, http.StatusAccepted, nil)
//	}
//}

//
//func (i *Api) GetAll(ctx context.Context) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		companies, err := i.companiesService.GetAll(ctx)
//		if err != nil {
//			i.Error(w, http.StatusInternalServerError, err)
//			return
//		}
//
//		i.Respond(w, http.StatusOK, bot.Response{Data: companies})
//	}
//}
//
//func (i *Api) Update(ctx context.Context) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var company model.Company
//		err := i.Decode(r, &company)
//		if err != nil {
//			i.Error(w, http.StatusBadRequest, errors.Join(bot.ErrDecodeBody, err))
//			return
//		}
//
//		err = i.companiesService.Update(ctx, &company)
//		if err != nil {
//			i.Error(w, http.StatusInternalServerError, err)
//			return
//		}
//
//		i.Respond(w, http.StatusAccepted, nil)
//	}
//}
