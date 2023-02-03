package clockface

import (
	"math"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestSecondHandPoint(t *testing.T) {
	testcases := []struct {
		time time.Time
		want Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 15), Point{1, 0}},
	}

	for _, tt := range testcases {
		t.Run(testName(tt.time), func(t *testing.T) {
			got := secondHandPoint(tt.time)
			if cmp.Equal(got, tt.want) {
				t.Errorf("secondHandPoint(%v) = %+v, want %+v", tt.time, got, tt.want)
			}
		})
	}
}

func Test_secondInRadiant(t *testing.T) {
	tests := []struct {
		name    string
		seconds int
		want    float64
	}{
		{"turn 0s into 0π", 0, 0},
		{"turn 30s into π", 30, math.Pi},
		{"turn 45s into 3/2π", 45, 3 * (math.Pi / 2)},
		{"turn 7s into 7/60π", 7, 7 * (math.Pi / 30)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := secondInRadiant(tt.seconds)
			if got != tt.want {
				t.Errorf("secondInRadiant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func simpleTime(hour, minute, second int) time.Time {
	return time.Date(666, time.December, 25, hour, minute, second, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
