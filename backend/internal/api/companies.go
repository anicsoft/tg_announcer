package api

import (
	"errors"
	"fmt"
	apiModel "tg_announcer/internal/api/model"
	"tg_announcer/internal/model"

	"github.com/gin-gonic/gin"
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
func (a *BaseApi) AddCompany(ctx *gin.Context) {
	company := model.NewCompany()

	err := ctx.ShouldBind(&company)
	if err != nil {
		StatusBadRequest(ctx, errors.Join(ErrDecodeBody, err))
		return
	}

	id, err := a.companiesService.Create(ctx, company)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusCreated(ctx, apiModel.AddCompanyResponse{ID: id})
}

// GetCompanyByID godoc
//
//	@Summary		Returns company by ID
//	@Description	Only for admins
//	@Tags			companies
//	@Produce		json
//	@Param			id	path		int	true	"request body"
//	@Success		200	{object}	model.Company
//	@Failure		400	{object}	HttpError	"failed to decode body"
//	@Failure		500	{object}	HttpError	"internal error"
//	@Router			/companies/{id} [get]
func (a *BaseApi) GetCompanyByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		StatusBadRequest(ctx, fmt.Errorf("empty id"))
		return
	}
	company, err := a.companiesService.GetByID(ctx, id)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusOK(ctx, company)
}

/*UpdateCompany(ctx *gin.Context)
DeleteCompany(ctx *gin.Context)
ListCompanies(ctx *gin.Context)*/

// UpdateCompany godoc
//
//	@Summary		Update a company
//	@Description	Update an existing company's information
//	@Tags			companies
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string			true	"Company ID"
//	@Param			company	body	model.Company	true	"Company data"
//	@Success		202
//	@Failure		400		{object}	HttpError	"failed to decode body or empty id"
//	@Failure		500		{object}	HttpError	"internal error"
//	@Router			/companies/{id} [patch]
func (a *BaseApi) UpdateCompany(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		StatusBadRequest(ctx, fmt.Errorf("empty id"))
		return
	}
	company := model.NewCompany()
	err := ctx.ShouldBind(&company)
	if err != nil {
		StatusBadRequest(ctx, errors.Join(ErrDecodeBody, err))
		return
	}
	company.ID = id
	err = a.companiesService.Update(ctx, company)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusAccepted(ctx)
}

// DeleteCompany godoc
//
//	@Summary		Delete a company
//	@Description	Delete a company by ID
//	@Tags			companies
//	@Produce		json
//	@Param			id	path	string	true	"Company ID"
//	@Success		202
//	@Failure		400	{object}	HttpError	"failed to decode body or empty id"
//	@Failure		500	{object}	HttpError	"internal error"
//	@Router			/companies/{id} [delete]
func (a *BaseApi) DeleteCompany(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		StatusBadRequest(ctx, fmt.Errorf("empty id"))
		return
	}
	err := a.companiesService.Delete(ctx, id)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusAccepted(ctx)
}

// ListCompanies godoc
//
//	@Summary		List all companies
//	@Description	Get a list of all companies
//	@Tags			companies
//	@Produce		json
//	@Success		200	{object}	model.Company
//	@Failure		400	{object}	HttpError	"failed to decode body"
//	@Failure		500	{object}	HttpError	"internal error"
//	@Router			/companies [get]
func (a *BaseApi) ListCompanies(ctx *gin.Context) {
	companies, err := a.companiesService.GetAll(ctx)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusOK(ctx, companies)
}
