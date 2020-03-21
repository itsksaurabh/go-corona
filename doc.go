/*
Package gocorna go-corona is a Golang client library for accessing global coronavirus (COVID-19, SARS-CoV-2) outbreak data.
It consumes data from Coronavirus Tracker API.

You can read the API server documentation at https://github.com/ExpDev07/coronavirus-tracker-api.

Usage:

create a new instance of Client, then use the various methods on the client to access different parts of the API.
For demonstration:
  package main
  import (
	"context"
	"fmt"
	"log"

   	"github.com/itsksaurabh/go-corona"
 )

  func main() {
  // client for accessing different endpoints of the API
	client := gocorona.Client{}
	ctx := context.Background()

  // GetLatestData returns total amonut confirmed cases, deaths, and recoveries.
	data, err := client.GetLatestData(ctx)
	if err != nil {
		log.Fatal("request failed:", err)
	}
	fmt.Println(data)
  }

Notes:
* Using the  [https://godoc.org/context](https://godoc.org/context) package for passing context.
* Look at tests(*_test.go) files for more sample usage.

*/
package gocorona
