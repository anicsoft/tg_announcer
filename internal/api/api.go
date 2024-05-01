package api

import (
	"anik/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Api interface {
	AddCompany(ctx context.Context) http.HandlerFunc
	AddAnnouncement(ctx context.Context) http.HandlerFunc
	AddOfferCategory(ctx context.Context) http.HandlerFunc
	AddBusinessCategory(ctx context.Context) http.HandlerFunc
	OfferCategories(ctx context.Context) http.HandlerFunc
	BusinessCategories(ctx context.Context) http.HandlerFunc
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
}

func New(
	companiesService service.CompaniesService,
	announcementService service.AnnouncementService,
	categoriesService service.CategoriesService,
) Api {
	return &BaseApi{
		companiesService:    companiesService,
		announcementService: announcementService,
		categoriesService:   categoriesService,
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
