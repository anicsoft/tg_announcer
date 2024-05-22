package api

import (
	apiModel "anik/internal/api/model"
	"anik/internal/model"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func (a *BaseApi) AddUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
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
func (a *BaseApi) Update(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req apiModel.UpdateUserRequest
		err := a.Decode(r, &req)
		if err != nil {
			a.Error(w, http.StatusBadRequest, err)
			return
		}

		user, err := a.userService.GetByID(ctx, req.UserID)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
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
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		a.Respond(w, http.StatusAccepted, nil)
	}
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
func (a *BaseApi) GetUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, _ := strconv.Atoi(idStr)
		user, err := a.userService.GetByID(ctx, id)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		a.Respond(w, http.StatusOK, user)
	}
}

func (a *BaseApi) Notify(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userFromRequest := model.NewUser()
		err := a.Decode(r, &userFromRequest)
		if err != nil {
			a.Error(w, http.StatusBadRequest, err)
			return
		}

		if !a.handleUserExistence(ctx, w, userFromRequest) {
			return
		}

		a.Respond(w, http.StatusOK, nil)
	}
}

func (a *BaseApi) handleUserExistence(ctx context.Context, w http.ResponseWriter, userFromRequest *model.User) bool {
	exists, _ := a.userService.Exists(ctx, userFromRequest.Id)
	//if err != nil {
	//	a.Error(w, http.StatusInternalServerError, err)
	//	return false
	//}

	if !exists {
		return a.handleNewUser(ctx, w, userFromRequest)
	}

	return a.handleExistingUser(ctx, w, userFromRequest)
}

func (a *BaseApi) handleNewUser(ctx context.Context, w http.ResponseWriter, userFromRequest *model.User) bool {
	id, err := a.userService.AddUser(ctx, userFromRequest)
	if err != nil {
		a.Error(w, http.StatusInternalServerError, err)
		return false
	}

	a.Respond(w, http.StatusOK, Response{Data: fmt.Sprintf("user added to db, id: %d", id)})
	return true
}

func (a *BaseApi) handleExistingUser(ctx context.Context, w http.ResponseWriter, userFromRequest *model.User) bool {
	userFromDB, err := a.userService.GetByID(ctx, userFromRequest.Id)
	if err != nil {
		a.Error(w, http.StatusInternalServerError, err)
		return false
	}
	log.Println("userfromdb:", userFromDB)
	if a.shouldUpdateUser(userFromDB, userFromRequest) {
		userFromRequest.UserType = userFromDB.UserType
		userFromRequest.CompanyId = userFromDB.CompanyId
		err := a.userService.Update(ctx, userFromRequest)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, fmt.Errorf("failed to update user: %w", err))
			return false
		}
	}

	return true
}

func (a *BaseApi) shouldUpdateUser(userFromDB, userFromRequest *model.User) bool {
	return userFromDB.Latitude != userFromRequest.Latitude || userFromDB.Longitude == userFromRequest.Longitude
}
