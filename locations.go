package gocorona

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// Coordinates hols coordinates of a location
type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// CaseCountWithTimestamp holds timestamp with CaseCount
type CaseCountWithTimestamp struct {
	Timestamp time.Time
	CaseCount int
}

// Timeline holds list of Case Counts with timestamp
type Timeline struct {
	Data []CaseCountWithTimestamp
}

// LatestWithTimeline struct holds latest count with timelines
type LatestWithTimeline struct {
	Latest   int      `json:"latest"`
	Timeline Timeline `json:"timeline"`
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// coverts json keys from string to unix timestamp and
// it's values as case count.
func (t *LatestWithTimeline) UnmarshalJSON(data []byte) error {
	type Alias LatestWithTimeline
	aux := struct {
		Timeline map[string]int `json:"timeline"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var temp CaseCountWithTimestamp
	for k, v := range aux.Timeline {
		timestamp, err := time.Parse(time.RFC3339, k)
		if err != nil {
			return err
		}

		temp.Timestamp = timestamp
		temp.CaseCount = v
		t.Timeline.Data = append(t.Timeline.Data, temp)
	}
	return nil
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

// GetAllLocationData returns all cases from all locations.
// You can Exclude/Include timelines. Timelines are excluded by default.
func (c Client) GetAllLocationData(ctx context.Context, timelines bool) (data Locations, err error) {
	t := "0"
	if timelines {
		t = "1"
	}

	endpoint := "/locations" + "?timelines=" + t
	if err = c.makeGetRequest(ctx, endpoint, &data); err != nil {
		return Locations{}, err
	}
	return data, nil
}

// GetDataByCountryCode returns all cases from different locations
// of a country by it's Country Code.
// Check alpha-2 country codes here: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
func (c Client) GetDataByCountryCode(ctx context.Context, countryCode string, timelines bool) (data Locations, err error) {
	if countryCode == "" {
		return Locations{}, errors.New("country code required")
	}

	t := "0"
	if timelines {
		t = "1"
	}

	endpoint := "/locations?country_code=" + countryCode + "&timelines=" + t
	if err = c.makeGetRequest(ctx, endpoint, &data); err != nil {
		return Locations{}, err
	}
	return data, nil
}

// LocationData holds data of a particular location
type LocationData struct {
	Location Location `json:"location"`
}

// GetDataByLocationID returns data of a specific location by it's ID.
// You can Exclude/Include timelines. Timelines are excluded by default.
// You can Exclude/Include timelines. Timelines are excluded by default.
func (c Client) GetDataByLocationID(ctx context.Context, id int, timelines bool) (data LocationData, err error) {
	t := "0"
	if timelines {
		t = "1"
	}

	endpoint := "/locations/" + strconv.Itoa(id) + "?timelines=" + t
	if err = c.makeGetRequest(ctx, endpoint, &data); err != nil {
		return LocationData{}, err
	}

	return data, nil
}
