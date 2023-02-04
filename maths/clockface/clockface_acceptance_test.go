package clockface_test

import (
	"bytes"
	"encoding/xml"
	"math"
	"testing"
	"time"

	"github.com/huytranpk2000/learn-go-with-tests/maths/clockface"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/google/go-cmp/cmp"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Lines   []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	Text string  `xml:",chardata"`
	X1   float64 `xml:"x1,attr"`
	Y1   float64 `xml:"y1,attr"`
	X2   float64 `xml:"x2,attr"`
	Y2   float64 `xml:"y2,attr"`
}

func TestSVGWriterSecondHand(t *testing.T) {
	testcases := []struct {
		time time.Time
		want Line
	}{
		{time: simpleTime(0, 0, 0), want: secondHandLine(150, 60)},
		{time: simpleTime(0, 0, 30), want: secondHandLine(150, 240)},
		{time: simpleTime(0, 0, 10), want: secondHandLine(150+math.Sqrt(3)*90/2, 150-45)},
	}
	for _, tt := range testcases {
		t.Run(testName(tt.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, tt.time)

			svg := SVG{}
			err := xml.Unmarshal(b.Bytes(), &svg)
			if err != nil {
				t.Fatalf("Didn't expected an error: %v", err)
			}

			if !containLine(tt.want, svg.Lines) {
				t.Errorf("Expected to find the second hand line %+v, in SVG lines %+v", tt.want, svg.Lines)
			}
		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	testcases := []struct {
		time time.Time
		want Line
	}{
		{time: simpleTime(0, 0, 0), want: minuteHandLine(150, 70)},
		{time: simpleTime(0, 30, 0), want: minuteHandLine(150, 230)},
		{time: simpleTime(0, 10, 0), want: minuteHandLine(150+math.Sqrt(3)*80/2, 150-40)},
	}
	for _, tt := range testcases {
		t.Run(testName(tt.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, tt.time)

			svg := SVG{}
			err := xml.Unmarshal(b.Bytes(), &svg)
			if err != nil {
				t.Fatalf("Didn't expected an error: %v", err)
			}

			if !containLine(tt.want, svg.Lines) {
				t.Errorf("Expected to find the minute hand line %+v, in SVG lines %+v", tt.want, svg.Lines)
			}
		})
	}
}

func TestSVGWriterHour(t *testing.T) {
	testcases := []struct {
		time time.Time
		want Line
	}{
		{simpleTime(0, 0, 0), hourHandLine(150, 100)},
		{simpleTime(6, 0, 0), hourHandLine(150, 200)},
		{simpleTime(3, 0, 0), hourHandLine(200, 150)},
		{simpleTime(9, 0, 0), hourHandLine(100, 150)},
	}

	for _, tt := range testcases {
		t.Run(testName(tt.time), func(t *testing.T) {
			var buf bytes.Buffer
			clockface.SVGWriter(&buf, tt.time)

			var svg SVG
			err := xml.Unmarshal(buf.Bytes(), &svg)
			if err != nil {
				t.Error("Shouldn't have an error, but got: ", err)
			}

			if !containLine(tt.want, svg.Lines) {
				t.Errorf("Expected the hour hand line at %+v, in SVG lines: %+v", tt.want, svg.Lines)
			}
		})
	}
}

func containLine(line Line, lines []Line) bool {
	approxOpt := cmpopts.EquateApprox(0.000001, 0.001)
	for _, l := range lines {
		if cmp.Equal(l, line, approxOpt) {
			return true
		}
	}

	return false
}

func simpleTime(hour, minute, second int) time.Time {
	return time.Date(666, time.December, 25, hour, minute, second, 0, time.UTC)
}

func secondHandLine(x2, y2 float64) Line {
	return Line{
		Text: clockface.SecondHandText,
		X1:   150,
		Y1:   150,
		X2:   x2,
		Y2:   y2,
	}
}

func minuteHandLine(x2, y2 float64) Line {
	return Line{
		Text: clockface.MinuteHandText,
		X1:   150,
		Y1:   150,
		X2:   x2,
		Y2:   y2,
	}
}

func hourHandLine(x2, y2 float64) Line {
	return Line{
		Text: clockface.HourHandText,
		X1:   150,
		Y1:   150,
		X2:   x2,
		Y2:   y2,
	}
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
