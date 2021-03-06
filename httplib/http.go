package httplib

import (
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/bjma/gurl/handler"
)

type HttpRequest struct {
	req  *http.Request
	resp *http.Response
	url  *url.URL
	dump []byte
}

// Constructs a new HTTP request
func NewHttpRequest(uri, method string) *HttpRequest {
	u, err := url.Parse(uri)
	if err != nil {
		err = handler.NewError("unable to create request with URL: " + u.String())
		handler.HandleError(err)
	}
	r := http.Request{
		Method:     method,
		URL:        u,
		Header:     make(http.Header),
		Proto:      "HTTP/2.0",
		ProtoMajor: 2,
		ProtoMinor: 0,
	}
	var res *http.Response
	return &HttpRequest{url: u, req: &r, resp: res}
}

// HTTP GET
func Get(uri string) *HttpRequest {
	return NewHttpRequest(uri, http.MethodGet)
}

// HTTP PUT
func Put(uri string, d string) *HttpRequest {
	r := NewHttpRequest(uri, http.MethodPut)
	r.req.ContentLength = int64(len(d))
	r = Body(r, d)
	return r
}

// HTTP POST
func Post(uri string, d string) *HttpRequest {
	r := NewHttpRequest(uri, http.MethodPost)
	r.req.ContentLength = int64(len(d))
	r = Body(r, d)
	return r
}

func Head(uri string) *HttpRequest {
	return NewHttpRequest(uri, http.MethodHead)
}

// Issues an HTTP request and returns the response
func SendRequest(r *HttpRequest) (*http.Response, error) {
	dump, err := httputil.DumpRequest(r.req, true)
	if err != nil {
		handler.HandleError(err)
	}
	r.dump = dump
	client := &http.Client{}

	resp, err := client.Do(r.req)
	if err != nil {
		return nil, err
	}
	// set response field
	r.resp = resp
	return resp, err
}

// Getters
func Response(r *HttpRequest) *http.Response {
	resp, err := SendRequest(r)
	if err != nil {
		handler.HandleError(err)
	}
	return resp
}

func Dump(r *HttpRequest) []byte {
	return r.dump
}

// Adds content to request body (PUT, POST)
func Body(r *HttpRequest, d string) *HttpRequest {
	// https://stackoverflow.com/questions/33606330/golang-rewrite-http-request-body
	r.req.Body = ioutil.NopCloser(strings.NewReader(d))
	return r
}

func Method(r *HttpRequest) string {
	return r.req.Method
}

func Header(r *HttpRequest, key string) string {
	return r.req.Header.Get(key)
}

func ContentLength(r *HttpRequest) int64 {
	return r.req.ContentLength
}

// Setters
func SetHeader(r *HttpRequest, key, value string) {
	// https://pkg.go.dev/net/http#Header
	r.req.Header.Set(key, value)
}
