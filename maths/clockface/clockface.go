package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func secondHandPoint(t time.Time) Point {
	rad := secondInRadiant(t.Second())
	return radiantToPoint(rad)
}

func minuteHandPoint(t time.Time) Point {
	rad := minuteInRadiant(t.Minute())
	return radiantToPoint(rad)
}

func hourHandPoint(t time.Time) Point {
	rad := hourInRadiant(t.Hour())
	return radiantToPoint(rad)
}

func radiantToPoint(rad float64) Point {
	x := math.Sin(rad)
	y := math.Cos(rad)

	return Point{X: x, Y: y}
}

func secondInRadiant(second int) float64 {
	return math.Pi / (30 / float64(second))
}

func minuteInRadiant(minute int) float64 {
	return math.Pi / (30 / float64(minute))
}

func hourInRadiant(hour int) float64 {
	return math.Pi / (6 / float64(hour))
}
