package api

import (
	"anik/internal/model"
	"context"
	"net/http"
)

func (i *Implementation) Create(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		company := model.NewCompany()
		err := i.decode(r, &company)
		if err != nil {
			i.error(w, http.StatusBadRequest, err)
			return
		}

		create, err := i.companiesService.Create(ctx, company)
		if err != nil {
			i.error(w, http.StatusInternalServerError, err)
			return
		}

		i.respond(w, http.StatusCreated, Response{Data: create})
	}
}
