package gspec

import (
  "bytes"
)

type FakeBody struct {
  bytes.Buffer
}

func (f *FakeBody) Close() error {
  return nil
}
