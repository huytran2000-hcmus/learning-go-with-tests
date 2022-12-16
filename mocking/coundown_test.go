package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("print 3 to Go!", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Countdown(&buffer, &spyCountdownOperations{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spyOutputAndPrinter := spyCountdownOperations{}
		Countdown(&spyOutputAndPrinter, &spyOutputAndPrinter)
		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}
		got := spyOutputAndPrinter.Calls
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := spyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.sleepDuration != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.sleepDuration)
	}
}

type spyTime struct {
	sleepDuration time.Duration
}

func (s *spyTime) Sleep(duration time.Duration) {
	s.sleepDuration = duration
}

type spyCountdownOperations struct {
	Calls []string
}

func (s *spyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleepOperation)
}

func (s *spyCountdownOperations) Write(p []byte) (count int, err error) {
	s.Calls = append(s.Calls, writeOperation)
	return
}

const (
	writeOperation = "write"
	sleepOperation = "sleep"
)
