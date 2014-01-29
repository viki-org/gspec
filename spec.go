package gspec

import (
  "fmt"
  "reflect"
  "testing"
)

type S struct {
  t *testing.T
}

type SR struct {
  fatal bool
  t *testing.T
  expected interface{}
}

func New(t *testing.T) (*S) {
  return &S{t:t}
}

func (s *S) Expect(expected interface{}, garbage ...interface{}) (sr *SR) {
  return &SR{t: s.t, expected: expected,}
}

func (s *S) ExpectOrFatal(expected interface{}, garbage ...interface{}) (sr *SR) {
  return &SR{t: s.t, expected: expected, fatal: true}
}

func (sr *SR) ToEqual(actual interface{}) {
  if sr.expected != actual {
    sr.fail(fmt.Sprintf("expected %+v to equal %+v", sr.expected, actual))
  }
}

func (sr *SR) ToNotEqual(actual interface{}) {
  if sr.expected == actual {
    sr.fail(fmt.Sprintf("expected %+v to not equal %+v", sr.expected, actual))
  }
}

func (sr *SR) ToBeNil() {
  if sr.expected == nil { return }
  if reflect.ValueOf(sr.expected).IsNil() { return }
  sr.fail(fmt.Sprintf("expected %+v to be nil", sr.expected))
}

func (sr *SR) ToNotBeNil() {
  if sr.expected != nil { return }
  if !reflect.ValueOf(sr.expected).IsNil() { return }
  sr.fail(fmt.Sprintf("expected %+v to not be nil", sr.expected))
}

func (sr *SR) fail(msg string) {
  if sr.fatal { sr.t.Fatal(msg) } else { sr.t.Error(msg) }
}
