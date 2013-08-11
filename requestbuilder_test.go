package gspec

import (
  "testing"
)

func TestSetsTheHeader(t *testing.T) {
  spec := New(t)
  req := Request().WithHeader("content-type", "application/json").Req
  spec.Expect(req.Header.Get("Content-Type")).ToEqual("application/json")
}
