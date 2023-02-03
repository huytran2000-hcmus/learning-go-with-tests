package clockface

import (
	"math"
	"time"
)

const (
	secondHandLength = 90
	clockCentreX     = 150
	clockCentreY     = 150
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)

	p.X = p.X * secondHandLength
	p.Y = p.Y * secondHandLength

	p.X = clockCentreX + p.X
	p.Y = clockCentreY - p.Y

	return p
}

func secondHandPoint(t time.Time) Point {
	radiant := secondInRadiant(t.Second())
	x := math.Sin(radiant)
	y := math.Cos(radiant)

	return Point{X: x, Y: y}
}

func secondInRadiant(second int) float64 {
	return math.Pi / (30 / float64(second))
}
