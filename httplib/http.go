package httplib

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

var defaultHttpSetting = HttpSettings{"gurl", 30 * time.Second, true}

// https://blog.golang.org/http-tracing

type HttpSettings struct {
	UserAgent      string
	ConnectTimeout time.Duration
	DumpBody       bool
}

// https://developer.mozilla.org/en-US/docs/Web/HTTP/Messages
type HttpRequest struct {
	req     *http.Request
	resp    *http.Response
	setting HttpSettings // CLI settings
	url     *url.URL
	dump    []byte
}

// Constructs a new HTTP request message
func NewHttpRequest(uri, method string) *HttpRequest {
	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
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
	return &HttpRequest{url: u, setting: defaultHttpSetting, req: &r, resp: res}
}

func SendRequest(r *HttpRequest) (*http.Response, error) {
	dump, err := httputil.DumpRequest(r.req, r.setting.DumpBody)
	if err != nil {
		fmt.Println(err.Error())
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

// Need bettter name tbh
func GetResponse(r *HttpRequest) (*http.Response, error) {
	resp, err := SendRequest(r)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func Response(r *HttpRequest) *http.Response {
	return r.resp
}

func Dump(r *HttpRequest) []byte {
	return r.dump
}

func SetHost(r *HttpRequest, host string) {
	r.req.Host = host
}

// https://pkg.go.dev/net/http#Header
func SetHeader(r *HttpRequest, key, value string) {
	r.req.Header.Set(key, value)
}

// Adds content to request body (PUT, POST)
func SetBody(r *HttpRequest, data []byte) {
	//r.req.Body = data
}
