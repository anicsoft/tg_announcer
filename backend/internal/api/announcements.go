package api

import (
	"errors"
	apiModel "tg_announcer/internal/api/model"
	"tg_announcer/internal/model"

	"github.com/gin-gonic/gin"
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
func (a *BaseApi) AddAnnouncement(ctx *gin.Context) {
	announcement := model.NewAnnouncement()
	err := ctx.ShouldBind(&announcement)
	if err != nil {
		StatusBadRequest(ctx, errors.Join(ErrDecodeBody, err))
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
		StatusInternalServerError(ctx, err)
		return
	}

	StatusOK(ctx, apiModel.AddAnnouncementResponse{ID: id})
}

// GetAnnouncements godoc
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
func (a *BaseApi) GetAnnouncements(ctx *gin.Context) {
	var filter apiModel.Filter

	err := ctx.ShouldBind(&filter)
	if err != nil {
		StatusBadRequest(ctx, errors.Join(ErrDecodeBody, err))
		return
	}

	announcements, err := a.announcementService.GetAll(ctx, filter)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusOK(ctx, apiModel.AnnouncementResponse{
		Announcements: announcements,
	})
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
func (a *BaseApi) GetAnnouncement(ctx *gin.Context) {
	id := ctx.Param("id")
	announcement, err := a.announcementService.Get(ctx, id)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusOK(ctx, announcement)
}

func (a *BaseApi) UpdateAnnouncements(ctx *gin.Context) {

}
