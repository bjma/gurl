package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/bjma/gurl/filelib"
	"github.com/bjma/gurl/handler"
	"github.com/bjma/gurl/httplib"
	"github.com/bjma/gurl/utils"
)

// HTTP options
var (
	method  = flag.String("X", "GET", "HTTP method")
	URI     = flag.String("url", "", "HTTP request URL") // https://site.com, :/portNum, :, site.com
	headers = flag.String("H", "", "HTTP headers in key:value format. To append multiple headers use comma separators without whitespace")
)

// Data options
var (
	output = flag.String("o", "", "Write data to output; defaulted to stdout")
	data   = flag.String("d", "", "Read data to output; defaulted to stdin")
)

// CLI output options
var (
	silent  = flag.Bool("s", false, "Supress HTTP headers")
	verbose = flag.Bool("v", false, "Print stuff for debugging")
)

func main() {
	flag.Parse()
	// NOTE: Add a httplib to parse command line to search for URL + a URL parser
	if len(*URI) == 0 {
		err := handler.NewError("empty URL")
		handler.HandleError(err)
	}
	doHTTP(*URI, *method)
}

// Sets the request headers, send HTTP request
func doHTTP(url, method string) {
	var req *httplib.HttpRequest

	// Initialize request according to flags
	switch method {
	case "GET":
		req = httplib.Get(url)
	case "PUT":
		if len(*data) == 0 {
			// Or, set Content-Length = 0
			err := handler.NewError("no request body")
			handler.HandleError(err)
		}
		body := httplib.ParseRequestBody(*data)
		contentLen := utils.Int64ToStr(int64(len(*data)))
        req = httplib.Put(url, body)
        if filelib.GetFileExtension(*data) == "json" {
			httplib.SetHeader(req, "Content-Type", "application/json")
		}
		httplib.SetHeader(req, "Content-Length", contentLen)
	case "POST":
		if len(*data) == 0 {
			err := handler.NewError("no request body")
			handler.HandleError(err)
		}
		body := httplib.ParseRequestBody(*data)
		contentLen := utils.Int64ToStr(int64(len(*data)))
		req = httplib.Post(url, body)
		if filelib.GetFileExtension(*data) == "json" {
			httplib.SetHeader(req, "Content-Type", "application/json")
		}
		httplib.SetHeader(req, "Content-Length", contentLen)
	default:
		req = httplib.Get(url)
	}

	// Default headers
	httplib.SetHeader(req, "User-Agent", "gurl/1.0")
	httplib.SetHeader(req, "Accept", "*/*")
	if (httplib.Header(req, "Content-Type") == "") {
        httplib.SetHeader(req, "Content-Type", "text/plain")
    } 
	// Custom headers
	if len(*headers) > 0 {
		for _, header := range strings.Split(*headers, ",") {
			token := strings.Split(header, ":")
			k := token[0]
			v := token[1]
			httplib.SetHeader(req, k, v)
		}
	}

	// Issue HTTP request
	resp := httplib.Response(req)

	// Formatting
	reqHeader := httplib.Dump(req)
	respHeader := httplib.FormatResponseHeader(method, resp)
	respBody := httplib.FormatResponseBody(resp)

	if !*silent {
		fmt.Println(string(reqHeader))
		if httplib.Method(req) != "GET" {
			// Extra newline for response body
			fmt.Printf("\n")
		}
		fmt.Println(string(respHeader))
	}

	if len(*output) > 0 {
		path := filelib.ParseFile(*output)
		bytesWritten := filelib.WriteFile(path, respBody)

		if *verbose {
			fmt.Printf("Wrote %d bytes to %s:\n", bytesWritten, path)
		}
	} else {
		fmt.Println(string(respBody))
	}
}
