package clockface

import (
	"math"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/google/go-cmp/cmp"
)

func Test_secondHandPoint(t *testing.T) {
	testcases := []struct {
		time time.Time
		want Point
	}{
		{time: simpleTime(0, 0, 30), want: Point{0, -1}},
		{simpleTime(0, 0, 15), Point{1, 0}},
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, tt := range testcases {
		t.Run(testName(tt.time), func(t *testing.T) {
			got := secondHandPoint(tt.time)
			assertEqualPoints(t, got, tt.want)
		})
	}
}

func Test_minuteHandPoint(t *testing.T) {
	testcases := []struct {
		time time.Time
		want Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 15, 0), Point{1, 0}},
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, tt := range testcases {
		t.Run(testName(tt.time), func(t *testing.T) {
			got := minuteHandPoint(tt.time)
			assertEqualPoints(t, got, tt.want)
		})
	}
}

func Test_hourHandPoint(t *testing.T) {
	testcases := []struct {
		time time.Time
		want Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(12, 0, 0), Point{0, 1}},
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(3, 0, 0), Point{1, 0}},
		{simpleTime(9, 0, 0), Point{-1, 0}},
		{simpleTime(4, 0, 0), Point{math.Sqrt(3) / 2, -0.5}},
	}

	for _, tt := range testcases {
		t.Run(testName(tt.time), func(t *testing.T) {
			got := hourHandPoint(tt.time)
			assertEqualPoints(t, got, tt.want)
		})
	}
}

func simpleTime(hour, minute, second int) time.Time {
	return time.Date(666, time.December, 25, hour, minute, second, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func assertEqualPoints(t *testing.T, got, want Point) {
	approxOption := cmpopts.EquateApprox(1e-5, 0.001)
	result := cmp.Equal(got, want, approxOption)

	if !result {
		t.Errorf("got %#v, want %#v", got, want)
	}
}
