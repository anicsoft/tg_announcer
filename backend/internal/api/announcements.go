package api

import (
	apiModel "anik/internal/api/model"
	"anik/internal/model"
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
)

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
//	@Success		201				{object}	model.AddAnnouncementResponse
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

		/*data, _ := ctxInitData(r.Context())
		user, err := a.userService.GetByID(ctx, int(data.User.ID))
		if err != nil {
			a.Error(w, http.StatusNotFound, errors.Join(ErrUserNotFound, err))
			return
		}*/

		// TODO CHECK IF SUCH COMPANY EXISTS
		// a.companiesService.Get(ctx, announcement.CompanyID)

		/*if user.CompanyId == nil || *user.CompanyId != announcement.CompanyID {
			a.Error(w, http.StatusForbidden, ErrNotAllowed)
			return
		}*/

		id, err := a.announcementService.Create(ctx, announcement)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		a.Respond(w, http.StatusCreated, apiModel.AddAnnouncementResponse{ID: id})
	}
}

// Announcements godoc
//
//	@Summary		Returns list of announcements
//	@Description	Filter body is used to apply various filters to the announcements query.
//	@Description Categories: A list of category names to filter the announcements by (e.g., "Special Offer").
//	@Description PromoCode: Set to true to retrieve announcements with a promo code.
//	@Description Latitude and Longitude: The user's location, used to calculate and return the distance to the user in meters.
//	@Description SortBy: The field to sort the results by (e.g., "distance").
//	@Description SortOrder: The order of sorting, either "asc" for ascending or "desc" for descending.
//	@Description PageSize: The number of results to return per page.
//	@Description Offset: The number of results to skip before starting to return results.
//	@Tags			announcements
//	@Accept			json
//	@Produce		json
//	@Param			filter	body		model.Filter	true	"request body"
//	@Success		200		{object}	model.AnnouncementResponse
//	@Failure		500		{object}	HttpError	"internal error"
//	@Router			/announcements/filter [post]
func (a *BaseApi) Announcements(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var announcements []model.Announcement
		var filter apiModel.Filter

		err := a.Decode(r, &filter)
		if err != nil {
			a.Error(w, http.StatusBadRequest, errors.Join(ErrDecodeBody, err))
			return
		}

		announcements, err = a.announcementService.GetAll(ctx, filter)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		a.Respond(w, http.StatusOK, apiModel.AnnouncementResponse{
			Announcements: announcements,
		})
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
		id := chi.URLParam(r, "id")
		announcement, err := a.announcementService.Get(ctx, id)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		a.Respond(w, http.StatusOK, Response{Data: announcement})
	}
}

/*func (a *BaseApi) CompanyAnnouncements(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		if idParam == "" {
			a.Error(w, http.StatusBadRequest, fmt.Errorf("empty id"))
			return
		}

		id, _ := strconv.Atoi(idParam)
		offers, err := a.announcementService.GetCompanyAnnouncements(ctx, id)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		a.Respond(w, http.StatusOK, offers)
	}
}*/
