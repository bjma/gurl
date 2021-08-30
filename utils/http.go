package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
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

func FormatResponseBody(res *http.Response) []byte {
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	bodyStr := bodyBytes
	return bodyStr
}

// Files are prefixed by '@' symbol;
// else, treat as string data
func ParseRequestBody(data string) string {
	var reqBody string
	if strings.HasPrefix(data, "@") {
		reqBody = ParseFile(data)
	} else {
		reqBody = data
	}
	return reqBody
}