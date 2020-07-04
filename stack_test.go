package gocollectionstack

import (
	"testing"
)

func TestStack(t *testing.T) {
	var s Stack
	s.Push("world.")
	s.Push("Hello, ")
	var got string
	for s.Size() > 0 {
		got = got + s.Pop()
	}
	want := "Hello, world."
	if got != want {
		t.Errorf("Stack result = %q, should be = %q", got, want)
	}
}
