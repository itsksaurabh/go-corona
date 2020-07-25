package gocorona

import (
	"context"
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

	if err := c.makeGetRequest(ctx, endpoint, &data); err != nil {
		return LatestData{}, err
	}
	return data, nil
}
