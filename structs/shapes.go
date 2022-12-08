package shapes

import "math"

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t *Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

type Shape interface {
	Area() float64
}
