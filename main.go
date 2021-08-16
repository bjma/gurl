package main

import (
	"flag"

	"github.com/bjma/gurl/util"
)

var (
    method = flag.String("method", "GET", "HTTP method")
    URI    = flag.String("url", "", "HTTP request URL") // https://site.com, :/portNum, :, site.com 
)

func main() {
    flag.Parse()
    
    // Flag parsing
    // ...
    
    // Each separate case should parse req/response body, set headers, etc.
    util.DoHTTP(*URI, *method)
}
