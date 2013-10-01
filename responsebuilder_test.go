package gspec

import (
  "testing"
)

func TestSetsTheResponseStatusCode(t *testing.T) {
  spec := New(t)
  res := Response().Status(203)
  spec.Expect(res.Res.StatusCode).ToEqual(203)

}

func TestSetsTheResponseBodyFromAByteArray(t *testing.T) {
  spec := New(t)
  res := Response().Body([]byte{116, 104, 101, 32, 115, 112, 105, 99, 101}).Res
  actual := make([]byte, 30)
  n, _ := res.Body.Read(actual)
  spec.Expect(string(actual[0:n])).ToEqual("the spice")
  spec.Expect(res.ContentLength).ToEqual(int64(9))
}

func TestSetsTheResponseBodyFromAString(t *testing.T) {
  spec := New(t)
  res := Response().BodyString("must flow now").Res
  actual := make([]byte, 30)
  n, _ := res.Body.Read(actual)
  spec.Expect(string(actual[0:n])).ToEqual("must flow now")
  spec.Expect(res.ContentLength).ToEqual(int64(13))
}
