package gspec

import (
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

func (s *S) Expect(expected interface{}) (sr *SR) {
  return &SR{t: s.t, expected: expected,}
}

func (sr *SR) ToEqual(actual interface{}) {
  if sr.expected != actual {
    sr.t.Errorf("expected %+v to equal %+v", sr.expected, actual)
  }
}
