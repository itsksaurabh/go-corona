package gocorona

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

// Latest holds fields related to latest data
type Latest struct {
	Confirmed int `json:"confirmed"`
	Deaths    int `json:"deaths"`
	Recovered int `json:"recovered"`
}

// LatestData holds response from the endpoint /v2/latest
type LatestData struct {
	Data Latest `json:"latest"`
}

// GetLatestData returns total amonut confirmed cases, deaths, and recoveries.
func (c Client) GetLatestData(ctx context.Context) (data LatestData, err error) {
	endpoint := "/latest"

	r, err := http.NewRequest(http.MethodGet, DefaultBaseURL+endpoint, nil)
	if err != nil {
		return LatestData{}, errors.Wrap(err, "could not generate http request")
	}

	if err = c.Do(WithCtx(ctx, r), &data); err != nil {
		return LatestData{}, errors.Wrap(err, "request failed")
	}
	return data, nil
}
