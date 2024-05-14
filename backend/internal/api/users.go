package api

import (
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

func (a *BaseApi) Update(ctx context.Context) http.HandlerFunc {
	type request struct {
		UserType  string `json:"user_type"`
		CompanyId int    `json:"company_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request
		idParam := chi.URLParam(r, "id")
		if idParam == "" {
			a.Error(w, http.StatusBadRequest, fmt.Errorf("empty idParam"))
			return
		}
		id, _ := strconv.Atoi(idParam)

		err := a.Decode(r, &req)
		if err != nil {
			a.Error(w, http.StatusBadRequest, err)
			return
		}

		user, err := a.userService.GetByID(ctx, id)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}
		log.Println("user before update:", user)
		user.UserType = req.UserType
		if req.UserType == "business" {
			user.CompanyId = &req.CompanyId
		}

		err = a.userService.Update(ctx, user)
		if err != nil {
			a.Error(w, http.StatusInternalServerError, err)
			return
		}

		a.Respond(w, http.StatusOK, Response{Data: fmt.Sprintf("user %s has been updated", user.Username)})
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

		a.Respond(w, http.StatusOK, Response{Data: "userFromRequest exists"})
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
