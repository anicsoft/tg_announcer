package api

import (
	"context"
	"net/http"
)

func (i *Implementation) GetAll(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		companies, err := i.companiesService.GetAll(ctx)
		if err != nil {
			i.error(w, http.StatusInternalServerError, err)
			return
		}

		i.respond(w, http.StatusOK, Response{Data: companies})
	}
}
