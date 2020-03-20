package gocorona

import (
	"context"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

// Coordinates hols coordinates of a location
type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// LatestWithTimeline struct holds latest count with timelines
type LatestWithTimeline struct {
	Latest    int            `json:"latest"`
	Timelines map[string]int `json:"timeline"`
}

// Timelines holds latest data with timelines
type Timelines struct {
	Confirmed LatestWithTimeline `json:"confirmed"`
	Deaths    LatestWithTimeline `json:"deaths"`
	Recovered LatestWithTimeline `json:"recovered"`
}

// Location holds data of a location
type Location struct {
	Coordinates Coordinates `json:"coordinates"`
	Country     string      `json:"country"`
	CountryCode string      `json:"country_code"`
	ID          int         `json:"id"`
	Latest      Latest      `json:"latest"`
	Province    string      `json:"province"`
	Timelines   Timelines   `json:"timelines,omitempty"`
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

// GetDataByCountryCode returns all cases from different locations
// of a country by it's Country Code.
// Check alpha-2 country codes here: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
func (c Client) GetDataByCountryCode(ctx context.Context, countryCode string) (data Locations, err error) {
	if countryCode == "" {
		return Locations{}, errors.New("country code required")
	}

	endpoint := "/locations?country_code=" + countryCode

	r, err := http.NewRequest(http.MethodGet, DefaultBaseURL+endpoint, nil)
	if err != nil {
		return Locations{}, errors.Wrap(err, "could not generate http request")
	}

	if err = c.Do(WithCtx(ctx, r), &data); err != nil {
		return Locations{}, errors.Wrap(err, "request failed")
	}
	return data, nil
}

// LocationData holds data of a particular location
type LocationData struct {
	Location Location `json:"location"`
}

// GetDataByLocationID returns data of a specific location by it's ID.
// You can Exclude/Include timelines. Timelines are excluded by default.
func (c Client) GetDataByLocationID(ctx context.Context, id int, timelines bool) (data LocationData, err error) {
	t := "0"
	if timelines {
		t = "1"
	}
	endpoint := "/locations/" + strconv.Itoa(id) + "?timelines=" + t

	r, err := http.NewRequest(http.MethodGet, DefaultBaseURL+endpoint, nil)
	if err != nil {
		return LocationData{}, errors.Wrap(err, "could not generate http request")
	}

	if err = c.Do(WithCtx(ctx, r), &data); err != nil {
		return LocationData{}, errors.Wrap(err, "request failed")
	}
	return data, nil
}
