package api

import (
	"errors"
	"log"
	"strconv"
	apiModel "tg_announcer/internal/api/model"
	"tg_announcer/internal/model"

	"github.com/gin-gonic/gin"
)

func (a *BaseApi) AddUser(ctx *gin.Context) {

}

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
//	@Param			id	path	int	true	"user id"
//	@Success		200
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

	StatusCreated(ctx, user)
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
