package api

import (
	apiModel "anik/internal/api/model"
	"anik/internal/model"
	"context"
	"errors"
	"net/http"
	"strconv"
)

func (i *Implementation) Create(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		company := apiModel.NewCompany()
		err := i.decode(r, &company)
		if err != nil {
			i.error(w, http.StatusBadRequest, errors.Join(ErrDecodeBody, err))
			return
		}

		servModel, err := convToServiceModel(company)
		if err != nil {
			i.error(w, http.StatusBadRequest, errors.Join(ErrParseLoc, err))
			return
		}

		create, err := i.companiesService.Create(ctx, servModel)
		if err != nil {
			i.error(w, http.StatusInternalServerError, err)
			return
		}

		i.respond(w, http.StatusCreated, Response{Data: create})
	}
}

func convToServiceModel(apiModel *apiModel.Company) (*model.Company, error) {
	lat, err := parseStringToFloat(apiModel.Latitude)
	if err != nil {
		return nil, err
	}
	lon, err := parseStringToFloat(apiModel.Longitude)
	if err != nil {
		return nil, err
	}

	return &model.Company{
		Id:          apiModel.Id,
		Name:        apiModel.Name,
		Description: apiModel.Description,
		Address:     apiModel.Address,
		Latitude:    lat,
		Longitude:   lon,
		Who:         apiModel.Who,
		CreatedAt:   apiModel.CreatedAt,
	}, nil
}

func parseStringToFloat(str string) (float64, error) {
	floatValue, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return floatValue, nil
}
