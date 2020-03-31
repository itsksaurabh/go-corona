package gocorona_test

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	gocorona "github.com/itsksaurabh/go-corona"
	"github.com/pkg/errors"
)

var (
	testServer     *url.URL
	testDataDir    = "./testdata/"
	updateTestData = flag.Bool("update", false, "if set then update testdata else use saved testdata for testing.")
)

func TestMain(m *testing.M) {
	flag.Parse()

	// Run testServer for unit tests
	if !*updateTestData {
		server := httptest.NewServer(http.FileServer(http.Dir(testDataDir)))

		surl, err := url.Parse(server.URL)
		if err != nil {
			fmt.Println("testServer URL parse failed:", err)
			os.Exit(1)
		}
		testServer = surl

		defer server.Close()
	}

	os.Exit(m.Run())
	return
}

// testClient returns a gocorna.Client mainly for testing purposes.
// It behaves as a reverse proxy agent, it reads testfile and
// returns data as a response to requests made through it.
//
// If update flag is set, it will save the real data to testfile.
func testClient(t *testing.T) gocorona.Client {
	c := gocorona.Client{
		HTTP: http.DefaultClient,
	}

	if *updateTestData {
		c.HTTP.Transport = &saverTransport{t}
		return c
	}

	c.HTTP.Transport = &loaderTransport{t}
	return c
}

// saverTransport saves response body to testdata file
type saverTransport struct{ t *testing.T }

func (st saverTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		return resp, errors.Wrap(err, "request failed")
	}

	if resp.StatusCode != http.StatusOK {
		return resp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, errors.Wrap(err, "read body failed")
	}
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	err = ioutil.WriteFile(testDataDir+filename(st.t), body, 0644)
	return resp, errors.Wrap(err, "write file failed")
}

// loaderTransport loads response from testdata file
type loaderTransport struct{ t *testing.T }

func (lt loaderTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	return http.Get(testServer.String() + "/" + filename(lt.t))
}

func filename(t *testing.T) string {
	name := t.Name()
	if strings.Contains(name, "/") { // If a subtest
		name = name[strings.LastIndex(t.Name(), "/")+1:]
	}
	name = strings.TrimPrefix(name, "Test")
	return name + ".json"
}
