package clockface_test

import (
	"testing"
	"time"

	"github.com/huytranpk2000/learn-go-with-tests/maths/clockface"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1666, time.December, 25, 0, 0, 0, 0, time.UTC)
	got := clockface.SecondHand(tm)

	want := clockface.Point{X: 150, Y: 150 - 90}
	if got != want {
		t.Errorf("clockface.SecondHand(%v) = %+v, want %+v", tm, got, want)
	}
}

func TestSecondHandAt30s(t *testing.T) {
	tm := time.Date(1666, time.December, 25, 0, 0, 30, 0, time.UTC)
	got := clockface.SecondHand(tm)

	want := clockface.Point{X: 150, Y: 150 + 90}
	if got != want {
		t.Errorf("clockface.SecondHand(%v) = %+v, want %+v", tm, got, want)
	}
}
