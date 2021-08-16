package util

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/bjma/gurl/httplib"
)

// Sets the request header, send HTTP request,
// cooks output, and performs a bunch of stuff according to
// flags
// https://developer.mozilla.org/en-US/docs/Glossary/Request_header
func DoHTTP(url, method string) {
    // Initialize request according to flags
    req := httplib.NewHttpRequest(url, method)
    httplib.SetHeader(req, "User-Agent", "gurl/1.0")
    res, err := httplib.GetResponse(req)
    fmt.Println(string(httplib.GetDump(req)))
    if err != nil {
        panic(err)
    }
    fmt.Println(res)
}

func formatRequest(r *httplib.HttpRequest) string {
    var buf bytes.Buffer
    return buf.String()
}

// https://developer.mozilla.org/en-US/docs/Glossary/Response_header
func formatResponse(method string, res *http.Response) string {
	var buf bytes.Buffer
	return buf.String()
}