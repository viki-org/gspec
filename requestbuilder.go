package gspec

import (
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
    },
  }
  return rb.Url("/")
}

func (rb *RequestBuilder) Header(key, value string) *RequestBuilder {
  rb.Req.Header.Set(key, value)
  return rb
}

func (rb *RequestBuilder) Url(u string) *RequestBuilder {
  rb.Req.URL, _ = url.Parse(u)
  return rb
}
