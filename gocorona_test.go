package gocorona_test

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	gocorona "github.com/itsksaurabh/go-corona"
	"github.com/pkg/errors"
)

var (
	updateTestData = flag.Bool(
		"update",
		false,
		"if set then update testdata else use saved testdata for testing.",
	)
)

type roundTripFunc func(r *http.Request) (*http.Response, error)

// RoundTrip implements http.RoundTripper interface.
func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

// client returns a gocorna.Client mainly for testing purposes.
// It behaves as a reverse proxy agent, it reads testfile and
// returns data as a response to requests made through it.
//
// If update flag is set, it will save the real data to testfile.
func testClient(t *testing.T) gocorona.Client {
	c := gocorona.Client{
		HTTP: http.DefaultClient,
	}
	c.HTTP.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if flag.Parse(); *updateTestData {
			if err := writeTestData(t, r); err != nil {
				return nil, errors.Wrap(err, "write testdata failed")
			}
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(readTestData(t)),
		}, nil
	})

	return c
}

// returns filename for saving response JSON
func filename(t *testing.T) string {
	name := t.Name()
	if strings.Contains(name, "/") { // If a subtest
		name = name[strings.LastIndex(t.Name(), "/")+1:]
	}
	name = strings.TrimPrefix(name, "Test")
	return name + ".json"
}

// readTestData loads response from testdata files and returns
// a new Reader reading from []bytes
// It creates filename with the help on test name
func readTestData(t *testing.T) *bytes.Reader {
	raw, err := ioutil.ReadFile("./testdata/" + filename(t))
	if err != nil {
		t.Error(err)
	}
	return bytes.NewReader(raw)
}

// writeTestData writes testdata to files
// It creates filename with the help on test name
func writeTestData(t *testing.T, r *http.Request) error {
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		return errors.Wrap(err, "request failed")
	}

	// decodes resp.Body into raw
	var raw json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return errors.Wrap(err, "body decode failed")
	}

	err = ioutil.WriteFile("./testdata/"+filename(t), raw, 0644)
	if err != nil {
		return errors.Wrap(err, "write test data failed")
	}
	return nil
}
