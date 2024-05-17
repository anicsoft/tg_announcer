package api

import (
	"anik/internal/model"
	"context"
	"log"
	"net/http"
)

type AddAnnouncementResponse struct {
	ID int `json:"id"`
}

type AnnouncementsResponse struct {
	Announcements []model.Announcement `json:"announcements"`
}

// AddAnnouncement godoc
//
//	@Summary		Create an announcement
//	@Description	Add by json announcement
//	@Tags			announcements
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string					true	"tma initData"
//	@Param			announcement	body		model.AddAnnouncement	true	"request body"
//	@Success		201				{object}	AddAnnouncementResponse
//	@Failure		401				{object}	HttpError	"failed to decode body"
//	@Failure		404				{object}	HttpError	"user not found"
//	@Failure		403				{object}	HttpError	"not allowed"
//	@Router			/announcement [post]
func (a *BaseApi) AddAnnouncement(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		announcement := model.NewAnnouncement()
		err := a.Decode(r, &announcement)
		if err != nil {
			a.Error(w, http.StatusBadRequest, ErrDecodeBody)
			return
		}

		data, _ := ctxInitData(r.Context())
		user, err := a.userService.GetByID(ctx, int(data.User.ID))
		if err != nil {
			a.Error(w, http.StatusNotFound, ErrUserNotFound)
			return
		}

		if user.CompanyId == nil || *user.CompanyId != announcement.CompanyID {
			a.Error(w, http.StatusForbidden, ErrNotAllowed)
			return
		}

		id, err := a.announcementService.Create(ctx, announcement)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		a.Respond(w, http.StatusCreated, AddAnnouncementResponse{ID: id})
	}
}

// Announcements godoc
//
//	@Summary		Returns list of announcements
//	@Description
//	@Tags			announcements
//	@Accept			json
//	@Produce		json
//	@Success		200				{object}	AnnouncementsResponse
//	@Failure		500				{object}	HttpError	"internal error"
//	@Router			/announcement [get]
func (a *BaseApi) Announcements(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		var announcements []model.Announcement
		var err error
		log.Printf("%+v", query)
		if len(query) > 0 {
			announcements, err = a.announcementService.GetFiltered(ctx, query)
			if err != nil {
				a.Error(w, http.StatusInternalServerError, err)
				return
			}
		} else {
			announcements, err = a.announcementService.GetAll(ctx)
			if err != nil {
				a.Error(w, http.StatusInternalServerError, err)
				return
			}
		}

		a.Respond(w, http.StatusOK, AnnouncementsResponse{announcements})
	}
}
