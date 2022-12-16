package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Huy")

	got := buffer.String()
	want := "Hello, Huy"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
