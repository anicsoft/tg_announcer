package api

import (
	"anik/internal/model"
	"context"
	"net/http"
)

func (i *Implementation) Update(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var company model.Company
		err := i.decode(r, &company)
		if err != nil {
			i.error(w, http.StatusBadRequest, err)
			return
		}

		err = i.companiesService.Update(ctx, &company)
		if err != nil {
			i.error(w, http.StatusInternalServerError, err)
			return
		}

		i.respond(w, http.StatusAccepted, nil)
	}
}
