package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bjma/gurl/filelib"
	"github.com/bjma/gurl/httplib"
	"github.com/bjma/gurl/util"
)

var (
	method = flag.String("X", "GET", "HTTP method")
	URI    = flag.String("url", os.Args[1], "HTTP request URL") // https://site.com, :/portNum, :, site.com
)

// Data options
var (
	output = flag.String("o", "", "Write data to output; defaulted to stdout")
	data   = flag.String("d", "", "Read data to output; defaulted to stdin")
)

// CLI output options
var (
	silent  = flag.Bool("s", false, "Supress HTTP headers")
	verbose = flag.Bool("v", true, "Print stuff for debugging")
)

func main() {
	flag.Parse()
	// Flag parsing
	// ...
	// Each separate case should parse req/response body, set headers, etc.
	execHTTP(*URI, *method)
}

// Sets the request headers, send HTTP request
// https://pkg.go.dev/net/http#Header
// https://developer.mozilla.org/en-US/docs/Glossary/Request_header
func execHTTP(url, method string) {
	// Initialize request according to flags
	req := httplib.NewHttpRequest(url, method)
	httplib.SetHeader(req, "Accept", "*/*")
    httplib.SetHeader(req, "User-Agent", "gurl/1.0")

	resp, err := httplib.GetResponse(req)
	if err != nil {
		panic(err)
	}
	reqHeader := httplib.Dump(req)
	respHeader := util.FormatResponseHeader(method, resp)
	respBody := util.FormatResponseBody(resp)
	// Print headers if output not supressed
	if !*silent {
		fmt.Println(string(reqHeader))
		fmt.Println(string(respHeader))
	}
	// Parse output param
	if len(*output) > 0 {
		filePath := util.ParseFile(*output)
		bytesWritten := filelib.WriteFile(filePath, respBody)
		if *verbose {
			fmt.Printf("Wrote %d bytes to ./tmp/%s:\n", bytesWritten, util.ParsePath(filePath))
		}
	} else {
		fmt.Println(string(respBody))
	}
}
