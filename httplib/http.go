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
    url *url.URL
    setting HttpSettings
    req *http.Request
    res *http.Response
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
        Proto: "HTTP/1.1",
        ProtoMajor: 1,
        ProtoMinor: 1,
    }
    var resp *http.Response
    return &HttpRequest{url: u, setting: defaultHttpSetting, req: &r, res: resp}
}

func SendRequest(r *HttpRequest) (*http.Response, error) {
    dump, err := httputil.DumpRequest(r.req, r.setting.DumpBody)
    if err != nil {
        fmt.Println(err.Error())
    }
    r.dump = dump
    client := &http.Client{}
    res, err := client.Do(r.req)
    if err != nil {
        return nil, err
    }
    // set response field 
    r.res = res
    return res, err
}

func GetResponse(r *HttpRequest) (*http.Response, error) {
    res, err := SendRequest(r)
    if err != nil {
        return nil, err
    }
    return res, err
}

func GetDump(r *HttpRequest) []byte {
    return r.dump
}



func SetHost(r *HttpRequest, host string) {
    r.req.Host = host
}

// https://pkg.go.dev/net/http#Header
func SetHeader(r *HttpRequest, key, value string) {
    r.req.Header.Set(key, value)
}
