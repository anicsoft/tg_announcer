package api

import (
	"anik/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Api interface {
	Notify(ctx context.Context) http.HandlerFunc
	Update(ctx context.Context) http.HandlerFunc
	AddUser(ctx context.Context) http.HandlerFunc
	AddCompany(ctx context.Context) http.HandlerFunc
	AddAnnouncement(ctx context.Context) http.HandlerFunc
	AddOfferCategory(ctx context.Context) http.HandlerFunc
	AddBusinessCategory(ctx context.Context) http.HandlerFunc
	OfferCategories(ctx context.Context) http.HandlerFunc
	BusinessCategories(ctx context.Context) http.HandlerFunc
	Announcements(ctx context.Context) http.HandlerFunc
}

type Response struct {
	Data interface{} `json:"data"`
}

type Err struct {
	Error interface{} `json:"error"`
}

type BaseApi struct {
	companiesService    service.CompaniesService
	announcementService service.AnnouncementService
	categoriesService   service.CategoriesService
	userService         service.UsersService
}

func New(
	companiesService service.CompaniesService,
	announcementService service.AnnouncementService,
	categoriesService service.CategoriesService,
	userService service.UsersService,
) Api {
	return &BaseApi{
		companiesService:    companiesService,
		announcementService: announcementService,
		categoriesService:   categoriesService,
		userService:         userService,
	}
}

func (a *BaseApi) Error(w http.ResponseWriter, code int, err error) {
	a.Respond(w, code, Err{err.Error()})
}

func (a *BaseApi) Respond(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (a *BaseApi) Decode(r *http.Request, data interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return fmt.Errorf("decode json: %w", err)
	}
	return nil
}
