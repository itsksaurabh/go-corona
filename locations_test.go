package gocorona_test

import (
	"context"
	"testing"

	gocorona "github.com/itsksaurabh/go-corona"
)

func TestGetAllLocationData(t *testing.T) {
	client := gocorona.Client{}
	ctx := context.Background()

	_, err := client.GetAllLocationData(ctx)
	if err != nil {
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestGetDataByCountryCode(t *testing.T) {
	client := gocorona.Client{}
	ctx := context.Background()
	countryCode := "US"
	_, err := client.GetDataByCountryCode(ctx, countryCode)
	if err != nil {
		if err != nil {
			t.Fatal(err)
		}
	}
}
