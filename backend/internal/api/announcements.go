package api

import (
	apiModel "anik/internal/api/model"
	"anik/internal/model"
	"context"
	"errors"
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
//
//	@Description	Only users with a "business" user_type can access this endpoint. The company_id in the request must match the company_id of the user making the request.
//
//	@Tags			announcements
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string					true	"tma initData"
//	@Param			announcement	body		model.AddAnnouncement	true	"request body"
//	@Success		201				{object}	AddAnnouncementResponse
//	@Failure		401				{object}	HttpError	"failed to decode body"
//	@Failure		404				{object}	HttpError	"user not found"
//	@Failure		403				{object}	HttpError	"not allowed"
//	@Router			/announcements [post]
func (a *BaseApi) AddAnnouncement(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		announcement := model.NewAnnouncement()
		err := a.Decode(r, &announcement)
		if err != nil {
			a.Error(w, http.StatusBadRequest, errors.Join(ErrDecodeBody, err))
			return
		}

		data, _ := ctxInitData(r.Context())
		user, err := a.userService.GetByID(ctx, int(data.User.ID))
		if err != nil {
			a.Error(w, http.StatusNotFound, errors.Join(ErrUserNotFound, err))
			return
		}

		// TODO CHECK IF SUCH COMPANY EXISTS
		// a.companiesService.GetByID(ctx, announcement.CompanyID)

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
//	@Summary	Returns list of announcements
//	@Description
//	@Tags		announcements
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	AnnouncementsResponse
//	@Failure	500	{object}	HttpError	"internal error"
//	@Router		/announcements [post]
func (a *BaseApi) Announcements(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var announcements []model.Announcement
		var filters apiModel.Filter

		err := a.Decode(r, &filters)
		if err != nil {
			a.Error(w, http.StatusBadRequest, errors.Join(ErrDecodeBody, err))
			return
		}

		announcements, err = a.announcementService.GetFiltered(ctx, filters)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		a.Respond(w, http.StatusOK, AnnouncementsResponse{announcements})
	}
}

// GetAnnouncement godoc
//
//	@Summary		Get announcement
//	@Description	Returns full announcement info.
//	@Tags			announcements
//	@Produce		json
//	@Param			id	path	int	true	"announcement id"
//	@Success		200
//	@Failure		500	{object}	HttpError	"internal error"
//	@Router			/announcements/{id} [get]
func (a *BaseApi) GetAnnouncement(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		strId := chi.URLParam(r, "id")
		id, _ := strconv.Atoi(strId)
		announcement, err := a.announcementService.Get(ctx, id)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		a.Respond(w, http.StatusOK, announcement)
	}
}
