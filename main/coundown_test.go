package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	spySleeper := spySleeper{}
	Countdown(&buffer, &spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	if spySleeper.Call != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %q", got)
	}
}
