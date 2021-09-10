package httplib

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/bjma/gurl/filelib"
	"github.com/bjma/gurl/handler"
)

// https://developer.mozilla.org/en-US/docs/Glossary/Response_header
func FormatResponseHeader(method string, resp *http.Response) []byte {
	var buf bytes.Buffer
	// Format response header
	buf.WriteString(resp.Proto + " " + resp.Status + "\n")
	for k, v := range resp.Header {
		buf.WriteString(k + ": " + strings.Join(v, "; ") + "\n")
	}
	return buf.Bytes()
}

func FormatResponseBody(resp *http.Response) []byte {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		handler.HandleError(err)
	}
	bodyStr := bodyBytes
	return bodyStr
}

// Files are prefixed by '@' symbol;
// else, treat as string data
func ParseRequestBody(data string) string {
	var reqBody string
	if strings.HasPrefix(data, "@") {
		f := filelib.ParseFile(data)
		// Read data and return as request body
		reqBody = string(filelib.ReadFile(f))
	} else {
		reqBody = data
	}
	return reqBody
}
