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
