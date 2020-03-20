package gocorona

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

// Coordinates hols coordinates of a location
type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// Location holds data of a location
type Location struct {
	Coordinates Coordinates `json:"coordinates"`
	Country     string      `json:"country"`
	CountryCode string      `json:"country_code"`
	ID          int         `json:"id"`
	Latest      Latest      `json:"latest"`
	Province    string      `json:"province"`
}

// Locations holds response from endpoint /v2/locations
type Locations struct {
	Locations []Location `json:"locations"`
}

// GetAllLocationData returns all cases from all locations
func (c Client) GetAllLocationData(ctx context.Context) (data Locations, err error) {
	endpoint := "/locations"

	r, err := http.NewRequest(http.MethodGet, DefaultBaseURL+endpoint, nil)
	if err != nil {
		return Locations{}, errors.Wrap(err, "could not generate http request")
	}

	if err = c.Do(WithCtx(ctx, r), &data); err != nil {
		return Locations{}, errors.Wrap(err, "request failed")
	}
	return data, nil
}
