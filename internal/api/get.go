package api

import (
	"context"
	"net/http"
)

func (i *Implementation) Get(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		company, err := i.companiesService.Get(ctx, id)
		if err != nil {
			i.error(w, http.StatusInternalServerError, err)
			return
		}

		i.respond(w, http.StatusOK, Response{Data: company})
	}
}
