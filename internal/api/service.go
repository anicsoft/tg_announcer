package api

import (
	"anik/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Data interface{} `json:"data"`
}

type Error struct {
	Error interface{} `json:"error"`
}

type Implementation struct {
	companiesService service.CompaniesService
}

func NewImplementation(companiesService service.CompaniesService) *Implementation {
	return &Implementation{companiesService: companiesService}
}

func (i *Implementation) error(w http.ResponseWriter, code int, err error) {
	i.respond(w, code, Error{err.Error()})
}

func (i *Implementation) respond(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (i *Implementation) decode(r *http.Request, data interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return fmt.Errorf("decode json: %w", err)
	}
	return nil
}
