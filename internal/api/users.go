package api

import (
	"anik/internal/model"
	"context"
	"fmt"
	"net/http"
)

func (a *BaseApi) Notify(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := model.NewUser()
		err := a.Decode(r, &user)
		if err != nil {
			a.Error(w, http.StatusBadRequest, err)
			return
		}

		exists, err := a.userService.Exists(ctx, user.Id)
		if err != nil {
			return
		}

		if !exists {
			id, err := a.userService.AddUser(ctx, user)
			if err != nil {
				a.Error(w, http.StatusInternalServerError, err)
				return
			}

			a.Respond(w, http.StatusOK, Response{Data: fmt.Sprintf("user added to db, id: %d", id)})
		} else {
			a.Respond(w, http.StatusOK, Response{Data: "user exists"})
		}
	}
}

func (a *BaseApi) AddUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
