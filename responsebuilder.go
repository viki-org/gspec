package gspec

import (
  "net/http"
)

type ResponseBuilder struct {
  Res *http.Response
}

func Response() *ResponseBuilder {
  return &ResponseBuilder{ Res: &http.Response { } }
}

func (r *ResponseBuilder) Status(s int) *ResponseBuilder {
  r.Res.StatusCode = s
  return r
}

func (r *ResponseBuilder) Body(data []byte) *ResponseBuilder {
  fb := new(FakeBody)
  s, _ := fb.Write(data)
  r.Res.Body = fb
  r.Res.ContentLength = int64(s)
  return r
}

func (r *ResponseBuilder) BodyString(b string) *ResponseBuilder {
  return r.Body([]byte(b))
}
