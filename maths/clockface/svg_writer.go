package clockface

import (
	"fmt"
	"io"
	"time"
)

const (
	SecondHandText = "Second Hand"
	MinuteHandText = "Minute Hand"
	HourHandText   = "Hour Hand"
)

const (
	clockCentreX     = 150
	clockCentreY     = 150
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
)

func SVGWriter(w io.Writer, tm time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)

	io.WriteString(w, secondHandTag(tm))
	io.WriteString(w, minuteHandTag(tm))
	io.WriteString(w, hourHandTag(tm))

	io.WriteString(w, svgEnd)
}

func secondHandTag(tm time.Time) string {
	p := SecondHand(tm)
	return fmt.Sprintf(
		`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:1px;">%s</line>`,
		p.X,
		p.Y,
		SecondHandText,
	)
}

func minuteHandTag(tm time.Time) string {
	p := MinuteHand(tm)
	return fmt.Sprintf(
		`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;">%s</line>`,
		p.X,
		p.Y,
		MinuteHandText,
	)
}

func hourHandTag(tm time.Time) string {
	p := HourHand(tm)
	return fmt.Sprintf(
		`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;">%s</line>`,
		p.X,
		p.Y,
		HourHandText,
	)
}

func SecondHand(tm time.Time) Point {
	p := secondHandPoint(tm)
	return makeHand(p, secondHandLength)
}

func MinuteHand(tm time.Time) Point {
	p := minuteHandPoint(tm)

	return makeHand(p, minuteHandLength)
}

func HourHand(tm time.Time) Point {
	p := hourHandPoint(tm)

	return makeHand(p, hourHandLength)
}

func makeHand(p Point, length float64) Point {
	p = Point{X: p.X * length, Y: p.Y * length}
	p = Point{X: clockCentreX + p.X, Y: clockCentreY - p.Y}

	return p
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
