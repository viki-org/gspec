package gspec

import (
  "bytes"
  "strings"
  "net/url"
  "net/http"
)

type RequestBuilder struct {
  Req *http.Request
}

type FakeBody struct {
  bytes.Buffer
}

func (f *FakeBody) Close() error {
  return nil
}

func Request() *RequestBuilder {
  rb := &RequestBuilder{
    Req: &http.Request{
      Method: "GET",
      Body: new(FakeBody),
      Header: make(http.Header),
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

func (rb *RequestBuilder) Body(data []byte) *RequestBuilder {
  fb := new(FakeBody)
  fb.Write(data)
  rb.Req.Body = fb
  return rb
}

func (rb *RequestBuilder) BodyString(b string) *RequestBuilder {
  return rb.Body([]byte(b))
}
