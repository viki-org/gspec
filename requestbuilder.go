package gspec

import (
  "net/http"
)

type RequestBuilder struct {
  Req *http.Request
}

func Request() *RequestBuilder {
  return &RequestBuilder{
    Req: &http.Request{
      Header: make(http.Header),
    },
  }
}

func (rb *RequestBuilder) WithHeader(key, value string) *RequestBuilder {
  rb.Req.Header.Set(key, value)
  return rb
}
