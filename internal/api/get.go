package api

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (i *Implementation) Get(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		company, err := i.companiesService.Get(ctx, id)
		if err != nil {
			i.error(w, http.StatusInternalServerError, err)
			return
		}

		i.respond(w, http.StatusOK, Response{Data: company})
	}
}
