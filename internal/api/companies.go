package api

import (
	"anik/internal/model"
	"context"
	"errors"
	"net/http"
)

func (a *BaseApi) AddCompany(ctx context.Context) http.HandlerFunc {
	type response struct {
		Id int `json:"id"`
	}

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
		resp := response{Id: id}
		a.Respond(w, http.StatusCreated, Response{Data: resp})
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

//func (i *Api) Get(ctx context.Context) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		id := chi.URLParam(r, "id")
//		if id == "" {
//			i.Error(w, http.StatusBadRequest, fmt.Errorf("empty id"))
//			return
//		}
//
//		company, err := i.companiesService.Get(ctx, id)
//		if err != nil {
//			i.Error(w, http.StatusInternalServerError, err)
//			return
//		}
//
//		i.Respond(w, http.StatusOK, api.Response{Data: company})
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
//		i.Respond(w, http.StatusOK, api.Response{Data: companies})
//	}
//}
//
//func (i *Api) Update(ctx context.Context) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var company model.Company
//		err := i.Decode(r, &company)
//		if err != nil {
//			i.Error(w, http.StatusBadRequest, errors.Join(api.ErrDecodeBody, err))
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
