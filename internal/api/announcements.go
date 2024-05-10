package api

import (
	"anik/internal/model"
	"context"
	"log"
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
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		a.Respond(w, http.StatusCreated, Response{Data: response{ID: id}})
	}
}

func (a *BaseApi) Announcements(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		log.Printf("%+v", query)
		if len(query) > 0 {
			announcements, err := a.announcementService.GetFiltered(ctx, query)
			if err != nil {
				a.Error(w, http.StatusInternalServerError, err)
				return
			}

			a.Respond(w, http.StatusOK, Response{Data: announcements})
		} else {
			announcements, err := a.announcementService.GetAll(ctx)
			if err != nil {
				a.Error(w, http.StatusInternalServerError, err)
				return
			}

			a.Respond(w, http.StatusOK, Response{Data: announcements})
		}
	}
}
