package api

import (
	"anik/internal/model"
	"context"
	"net/http"
)

func (a *BaseApi) AddAnnouncement(ctx context.Context) http.HandlerFunc {
	type response struct {
		ID int `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		announcement := model.NewAnnouncement()
		err := a.Decode(r, &announcement)
		if err != nil {
			a.Error(w, http.StatusBadRequest, err)
			return
		}

		id, err := a.announcementService.Create(ctx, announcement)
		if err != nil {
			return
		}

		a.Respond(w, http.StatusCreated, Response{Data: response{ID: id}})
	}
}
