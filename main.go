package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bjma/gurl/filelib"
	"github.com/bjma/gurl/httplib"
	"github.com/bjma/gurl/util"
)

// HTTP options
var (
	method  = flag.String("X", "GET", "HTTP method")
	URI     = flag.String("url", os.Args[1], "HTTP request URL") // https://site.com, :/portNum, :, site.com
	headers = flag.String("H", "", "HTTP headers in key:value format. To append multiple headers use ; as separators")
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
	execHTTP(*URI, *method)
}

// Sets the request headers, send HTTP request
func execHTTP(url, method string) {
	var req *httplib.HttpRequest

	// Initialize request according to flags
	switch method {
	case "PUT":
		if len(*data) == 0 {
			// Or, set Content-Length = 0
			log.Fatalln("ERR: No request body")
		}
		body := util.ParseRequestBody(*data)
		contentLen := strconv.FormatInt(int64(len(*data)), 10)
		req = httplib.NewPutRequest(url, body)
		httplib.SetHeader(req, "Content-Length", contentLen)
	default:
		req = httplib.NewGetRequest(url)
	}

	// Default headers
	httplib.SetHeader(req, "Accept", "*/*")
	httplib.SetHeader(req, "User-Agent", "gurl/1.0")
	// Custom headers
	if len(*headers) > 0 {
		for _, header := range strings.Split(*headers, ",") {
			token := strings.Split(header, ":")
			k := token[0]
			v := token[1]
			httplib.SetHeader(req, k, v)
		}
	}

	resp, err := httplib.GetResponse(req)
	if err != nil {
		panic(err)
	}
	reqHeader := httplib.Dump(req)
	respHeader := util.FormatResponseHeader(method, resp)
	respBody := util.FormatResponseBody(resp)

	if !*silent {
		fmt.Println(string(reqHeader))
		if httplib.Method(req) != "GET" {
			// Extra newline for response body
			fmt.Printf("\n")
		}
		fmt.Println(string(respHeader))
	}

	if len(*output) > 0 {
		// idk if we need synchronization rn
		wlock := make(chan int, 1)

		path := util.ParseFile(*output)
		go filelib.WriteFile(path, respBody, wlock)
		bytesWritten := <-wlock
		if *verbose {
			fmt.Printf("Wrote %d bytes to %s:\n", bytesWritten, path)
		}
	} else {
		fmt.Println(string(respBody))
	}
}
