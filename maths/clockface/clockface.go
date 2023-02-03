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
	radiant := secondInRadiant(t.Second())
	x := math.Sin(radiant)
	y := math.Cos(radiant)

	return Point{X: x, Y: y}
}

func secondInRadiant(second int) float64 {
	return math.Pi / (30 / float64(second))
}
