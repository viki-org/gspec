package gspec

import (
  "reflect"
  "testing"
)

type S struct {
  t *testing.T
}

type SR struct {
  t *testing.T
  expected interface{}
}

func New(t *testing.T) (*S) {
  return &S{t:t}
}

func (s *S) Expect(expected interface{}, garbage ...interface{}) (sr *SR) {
  return &SR{t: s.t, expected: expected,}
}

func (sr *SR) ToEqual(actual interface{}) {
  if sr.expected != actual {
    sr.t.Errorf("expected %+v to equal %+v", sr.expected, actual)
  }
}

func (sr *SR) ToBeNil() {
  if sr.expected == nil { return }
  if reflect.ValueOf(sr.expected).IsNil() { return }
  sr.t.Errorf("expected %+v to be nil", sr.expected)
}

func (sr *SR) ToNotBeNil() {
  if sr.expected != nil { return }
  if !reflect.ValueOf(sr.expected).IsNil() { return }
  sr.t.Errorf("expected %+v to not be nil", sr.expected)
}
