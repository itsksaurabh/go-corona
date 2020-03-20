package gocorona_test

import (
	"context"
	"testing"

	gocorona "github.com/itsksaurabh/go-corona"
)

func TestGetAllLocationData(t *testing.T) {
	client := gocorona.Client{}
	ctx := context.Background()

	tests := map[string]struct {
		wantTimeline bool
	}{
		"With Timeline":    {wantTimeline: true},
		"Without Timeline": {wantTimeline: false},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := client.GetAllLocationData(ctx, tt.wantTimeline)
			if err != nil {
				if err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

func TestGetDataByCountryCode(t *testing.T) {
	client := gocorona.Client{}
	ctx := context.Background()
	countryCode := "US"

	tests := map[string]struct {
		wantTimeline bool
	}{
		"With Timeline":    {wantTimeline: true},
		"Without Timeline": {wantTimeline: false},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := client.GetDataByCountryCode(ctx, countryCode, tt.wantTimeline)
			if err != nil {
				if err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

func TestGetDataByLocationID(t *testing.T) {
	client := gocorona.Client{}
	ctx := context.Background()
	locationID := 123

	tests := map[string]struct {
		wantTimeline bool
	}{
		"With Timeline":    {wantTimeline: true},
		"Without Timeline": {wantTimeline: false},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := client.GetDataByLocationID(ctx, locationID, tt.wantTimeline)
			if err != nil {
				if err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}
