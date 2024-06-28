package api

import (
	"errors"
	"log"
	"strconv"
	apiModel "tg_announcer/internal/api/model"
	"tg_announcer/internal/model"

	"github.com/gin-gonic/gin"
)

// Update godoc
//
//	@Summary		Update user
//	@Description	This endpoint is restricted to admin users only. It updates the user_type to either "business" or "user". If the user_type is set to "business", you must also provide the company_id that the user belongs to.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header	string					true	"tma initData"
//	@Param			announcement	body	model.UpdateUserRequest	true	"request body"
//	@Success		202
//	@Failure		400	{object}	HttpError	"failed to decode body"
//	@Failure		500	{object}	HttpError	"internal error"
//	@Router			/users [patch]
func (a *BaseApi) Update(ctx *gin.Context) {
	var req apiModel.UpdateUserRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		StatusBadRequest(ctx, errors.Join(ErrDecodeBody, err))
		return
	}

	user, err := a.userService.GetByID(ctx, req.UserID)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	user.UserType = req.UserType
	if req.UserType == "business" {
		user.CompanyId = &req.CompanyId
	} else {
		user.CompanyId = nil
	}

	err = a.userService.Update(ctx, user)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusAccepted(ctx)
}

// GetUser godoc
//
//	@Summary		Get user
//	@Description	Returns full user info.
//	@Tags			users
//	@Produce		json
//	@Param			id	path		int	true	"user id"
//	@Success		200	{object}	model.User
//	@Failure		500	{object}	HttpError	"internal error"
//	@Router			/users/{id} [get]
func (a *BaseApi) GetUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	user, err := a.userService.GetByID(ctx, id)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusOK(ctx, user)
}

// ListUsers godoc
//
//	@Summary		List users
//	@Description	Get a list of all users
//	@Tags			users
//	@Produce		json
//	@Success		200	{object}	HttpError
//	@Failure		500	{object}	HttpError
//	@Router			/users [get]
func (a *BaseApi) ListUsers(ctx *gin.Context) {
	users, err := a.userService.UserList(ctx)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusOK(ctx, apiModel.ListUserResponse{
		Users: users,
	})
}

// AddFavorite godoc
//
//	@Summary		Add favorite company
//	@Description	Add a company to the user's list of favorite companies
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"User ID"
//	@Param			body	body		model.AddFavoriteRequest	true	"Add Favorite Request"
//	@Success		200		{object}	nil
//	@Failure		400		{object}	HttpError
//	@Failure		500		{object}	HttpError
//	@Router			/users/{id}/favorites [post]
func (a *BaseApi) AddFavorite(ctx *gin.Context) {
	idStr := ctx.Param("id")
	userId, err := strconv.Atoi(idStr)
	if err != nil {
		StatusBadRequest(ctx, err)
		return
	}

	var req apiModel.AddFavoriteRequest
	if err = ctx.ShouldBindJSON(&req); err != nil {
		StatusBadRequest(ctx, err)
		return
	}

	if err = a.userService.AddFavorite(ctx, userId, req.CompanyID); err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusOK(ctx, nil)
}

// ListFavorites godoc
//
//	@Summary		List favorite companies
//	@Description	Get a list of favorite companies for a user
//	@Tags			users
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	model.FavoritesResponse
//	@Failure		400	{object}	HttpError
//	@Failure		500	{object}	HttpError
//	@Router			/users/{id}/favorites [get]
func (a *BaseApi) ListFavorites(ctx *gin.Context) {
	idStr := ctx.Param("id")
	userId, err := strconv.Atoi(idStr)
	if err != nil {
		StatusBadRequest(ctx, err)
		return
	}

	companies, err := a.userService.Favorites(ctx, userId)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusOK(ctx, apiModel.FavoritesResponse{Companies: companies})
}

// DeleteFavorite godoc
//
//	@Summary		Delete favorite company
//	@Description	Delete a company from the user's list of favorite companies
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int								true	"User ID"
//	@Param			body	body		model.DeleteFavoriteRequest	true	"Delete Favorite Request"
//	@Success		200		{object}	nil
//	@Failure		400		{object}	HttpError
//	@Failure		500		{object}	HttpError
//	@Router			/users/{id}/favorites [delete]
func (a *BaseApi) DeleteFavorite(ctx *gin.Context) {
	idStr := ctx.Param("id")
	userId, err := strconv.Atoi(idStr)
	if err != nil {
		StatusBadRequest(ctx, err)
		return
	}

	var req apiModel.DeleteFavoriteRequest
	if err = ctx.ShouldBindJSON(&req); err != nil {
		StatusBadRequest(ctx, err)
		return
	}

	if err = a.userService.DeleteFavorite(ctx, userId, req.CompanyID); err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusOK(ctx, nil)
}

func (a *BaseApi) Notify(ctx *gin.Context) {
	userFromRequest := model.NewUser()
	if err := ctx.ShouldBind(&userFromRequest); err != nil {
		StatusBadRequest(ctx, errors.Join(ErrDecodeBody, err))
		return
	}

	log.Println("userFromRequest", userFromRequest)
	if _, err := a.userService.AddUser(ctx, userFromRequest); err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusOK(ctx, nil)
}
