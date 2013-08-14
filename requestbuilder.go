package gspec

import (
  "strings"
  "net/url"
  "net/http"
)

type RequestBuilder struct {
  Req *http.Request
}

func Request() *RequestBuilder {
  rb := &RequestBuilder{
    Req: &http.Request{
      Header: make(http.Header),
      Method: "GET",
      RemoteAddr: "127.0.0.1:19001",
    },
  }
  return rb.Url("/")
}

func (rb *RequestBuilder) Header(key, value string) *RequestBuilder {
  rb.Req.Header.Set(key, value)
  return rb
}

func (rb *RequestBuilder) Method(method string) *RequestBuilder {
  rb.Req.Method = strings.ToUpper(method)
  return rb
}

func (rb *RequestBuilder) RemoteAddr(addr string) *RequestBuilder {
  rb.Req.RemoteAddr = addr
  return rb
}

func (rb *RequestBuilder) Url(u string) *RequestBuilder {
  rb.Req.URL, _ = url.Parse(u)
  return rb
}
