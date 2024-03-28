package api

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (i *Implementation) Delete(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		err := i.companiesService.Delete(ctx, id)
		if err != nil {
			i.error(w, http.StatusInternalServerError, err)
			return
		}

		i.respond(w, http.StatusAccepted, nil)
	}
}
