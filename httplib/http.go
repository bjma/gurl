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
    Request *http.Request
    Response *http.Response
    setting HttpSettings // CLI settings
    url *url.URL
    dump []byte
}

// Constructs a new HTTP request message
func NewHttpRequest(uri, method string) *HttpRequest {
    u, err := url.Parse(uri)
    if err != nil {
        panic(err)
    }
    r := http.Request{
        Method: method,
        URL: u,
        Header: make(http.Header),
        Proto: "HTTP/2.0",
        ProtoMajor: 2,
        ProtoMinor: 0,
    }
    var resp *http.Response
    return &HttpRequest{url: u, setting: defaultHttpSetting, Request: &r, Response: resp}
}

func SendRequest(r *HttpRequest) (*http.Response, error) {
    dump, err := httputil.DumpRequest(r.Request, r.setting.DumpBody)
    if err != nil {
        fmt.Println(err.Error())
    }
    r.dump = dump
    client := &http.Client{}
    resp, err := client.Do(r.Request)
    if err != nil {
        return nil, err
    }
    // set response field 
    r.Response = resp
    return resp, err
}

func GetResponse(r *HttpRequest) (*http.Response, error) {
    resp, err := SendRequest(r)
    if err != nil {
        return nil, err
    }
    return resp, err
}

func GetDump(r *HttpRequest) []byte {
    return r.dump
}


func SetHost(r *HttpRequest, host string) {
    r.Request.Host = host
}

// https://pkg.go.dev/net/http#Header
func SetHeader(r *HttpRequest, key, value string) {
    r.Request.Header.Set(key, value)
}
