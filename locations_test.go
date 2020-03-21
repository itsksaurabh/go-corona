package gocorona_test

import (
	"context"
	"testing"
)

func TestGetAllLocationData(t *testing.T) {
	ctx := context.Background()

	tests := map[string]struct {
		wantTimeline bool
	}{
		"With Timeline":    {wantTimeline: true},
		"Without Timeline": {wantTimeline: false},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := testClient(t).GetAllLocationData(ctx, tt.wantTimeline)
			if err != nil {
				if err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

func TestGetDataByCountryCode(t *testing.T) {
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
			_, err := testClient(t).GetDataByCountryCode(ctx, countryCode, tt.wantTimeline)
			if err != nil {
				if err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

func TestGetDataByLocationID(t *testing.T) {
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
			_, err := testClient(t).GetDataByLocationID(ctx, locationID, tt.wantTimeline)
			if err != nil {
				if err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}
