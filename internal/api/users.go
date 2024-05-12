package api

import (
	"context"
	"log"
	"net/http"
)

func (a *BaseApi) Notify(ctx context.Context) http.HandlerFunc {
	type request struct {
		Id           int     `json:"id"`
		FirstName    string  `json:"first_name"`
		LastName     string  `json:"last_name"`
		Username     string  `json:"username"`
		LanguageCode string  `json:"language_code"`
		Latitude     float32 `json:"latitude"`
		Longitude    float32 `json:"longitude"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		err := a.Decode(r, &req)
		if err != nil {
			a.Error(w, http.StatusBadRequest, err)
			return
		}

		log.Println("request: ", req)

	}
}

func (a *BaseApi) AddUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
