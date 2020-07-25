package gocorona_test

import (
	"context"
	"testing"
)

func TestGetLatestData(t *testing.T) {
	client := testClient(t)
	ctx := context.Background()

	_, err := client.GetLatestData(ctx)
	if err != nil {
		if err != nil {
			t.Fatal(err)
		}
	}
}
