package gspec

import (
  "testing"
)

func TestSetsTheHeader(t *testing.T) {
  spec := New(t)
  req := Request().Header("content-type", "application/json").Req
  spec.Expect(req.Header.Get("Content-Type")).ToEqual("application/json")
}

func TestSetsTheUrl(t *testing.T) {
  spec := New(t)
  req := Request().Url("/level.json?power=9001").Req
  spec.Expect(req.URL.Path).ToEqual("/level.json")
  spec.Expect(req.URL.RawQuery).ToEqual("power=9001")
}

func TestSetsTheBodyFromAByteArray(t *testing.T) {
  spec := New(t)
  req := Request().Body([]byte{116, 104, 101, 32, 115, 112, 105, 99, 101}).Req
  actual := make([]byte, 30)
  n, _ := req.Body.Read(actual)
  spec.Expect(string(actual[0:n])).ToEqual("the spice")
}

func TestSetsTheBodyFromAString(t *testing.T) {
  spec := New(t)
  req := Request().BodyString("must flow").Req
  actual := make([]byte, 30)
  n, _ := req.Body.Read(actual)
  spec.Expect(string(actual[0:n])).ToEqual("must flow")
}
