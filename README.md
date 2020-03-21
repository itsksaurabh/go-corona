
# go-corona

[![itsksaurabh](https://circleci.com/gh/itsksaurabh/go-corona.svg?style=shield)](https://circleci.com/gh/itsksaurabh/workflows/go-corona/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/itsksaurabh/go-corona)](https://goreportcard.com/report/github.com/itsksaurabh/go-corona)
[![GoDoc](https://godoc.org/github.com/itsksaurabh/go-corona?status.svg)](https://godoc.org/github.com/itsksaurabh/go-corona)
[![Built with Mage](https://magefile.org/badge.svg)](https://magefile.org)
[![MIT License](https://img.shields.io/github/license/itsksaurabh/go-corona?style=social)](https://github.com/itsksaurabh/go-corona/blob/master/LICENSE)

<img style="float:left;" width="200" src="./assets/logo.png"> 

#### go-corona is a [Golang](http://golang.org/) client library for accessing global coronavirus (COVID-19, SARS-CoV-2) outbreak data.

## API Documentation
It consumes data from [Coronavirus Tracker API](https://github.com/ExpDev07/coronavirus-tracker-api). You can read the API server documentation [here](https://github.com/ExpDev07/coronavirus-tracker-api).

#### Available data sources:

* **JHU** - https://github.com/CSSEGISandData/COVID-19 - Data repository operated by the Johns Hopkins University Center for Systems Science and Engineering (JHU CSSE). More data sources to be added later.

## Installation

Make sure you have set the environment variable $GOPATH

```bash
export GOPATH="path/to/your/go/folder"
```

Obtain the latest version of the  go-corona library with:

```bash
go get github.com/itsksaurabh/go-corona
```

Then, add the following to your Golang project:

```go
import (
	"github.com/itsksaurabh/go-corona"
)
```

## Usage
Package provides a client for accessing different endpoints of the API.
Create a new instance of Client, then use the various methods on the client to access different parts of the API.

For demonstration:
```go
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

  ```
Notes:
* Using the  [https://godoc.org/context](https://godoc.org/context) package for passing context.
* Look at tests(*_test.go) files for more sample usage.

## Error Handling

All errors generated at runtime will be returned to the calling client method. Any API request for which the API returns an error encoded in a JSON response will be parsed and returned by the client method as a Golang error struct. Lastly, it is important to note that for HTTP requests, if the response code returned is not '200 OK', an error will be returned to the client method detailing the response code that was received.

## Testing

In order to run the tests for this library, you will first need to install [Mage](https://magefile.org/) - A Make/rake-like dev tool using Go. You can install the dependency with the following command:

**Using GOPATH**

```bash
go get -u -d github.com/magefile/mage
cd $GOPATH/src/github.com/magefile/mage
go run bootstrap.go
```

**Using Go Modules**

```bash
git clone https://github.com/magefile/mage
cd mage
go run bootstrap.go
```
The mage binary will be created in your `$GOPATH/bin` directory.
You may also install a binary release from Mage's [releases](https://github.com/magefile/mage/releases) page.

Then run all tests by executing the following in your command line:
    
 	$ mage -v Test

**Updating Test Data**

You can update the test data inside `./testdata/` by enabling the following flag inside the file [gocorona_test.go](./gocorona_test.go#L16) and then perform testing. By default the flag is set to `false`.
```go
var (
	updateTestData = flag.Bool(
		"update",
		true,
		"if set then update testdata else use saved testdata for testing.",
	)
)
```
# Contributing
I welcome pull requests, bug fixes and issue reports. Before proposing a change, please discuss your change by raising an issue.

# Maintainer 

[Kumar Saurabh](https://in.linkedin.com/in/itsksaurabh)

## License

[MIT](LICENSE) © Kumar Saurabh
