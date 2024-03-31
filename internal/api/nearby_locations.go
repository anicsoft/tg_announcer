package api

import (
	apiModel "anik/internal/api/model"
	"anik/internal/model"
	"context"
	"errors"
	"net/http"
)

func (i *Implementation) NearbyLocations(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loc apiModel.Location
		err := i.decode(r, &loc)
		if err != nil {
			i.error(w, http.StatusBadRequest, errors.Join(ErrDecodeBody, err))
			return
		}

		parsedLoc, err := i.parseLocation(&loc)
		if err != nil {
			i.error(w, http.StatusBadRequest, errors.Join(ErrParseLoc, err))
			return
		}

		locations, err := i.companiesService.NearbyLocations(ctx, parsedLoc)
		if err != nil {
			return
		}

		i.respond(w, http.StatusOK, Response{Data: locations})
	}
}

func (i *Implementation) parseLocation(loc *apiModel.Location) (*model.Location, error) {
	lat, err := parseStringToFloat(loc.Latitude)
	if err != nil {
		return nil, err
	}

	lon, err := parseStringToFloat(loc.Longitude)
	if err != nil {
		return nil, err
	}

	return &model.Location{
		Latitude:  lat,
		Longitude: lon,
	}, nil
}
