package clockface

import (
	"fmt"
	"io"
	"time"
)

const (
	SecondHandText = "Second Hand"
)

const (
	secondHandLength = 90
	clockCentreX     = 150
	clockCentreY     = 150
)

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)

	p := SecondHand(t)
	io.WriteString(w, secondHandTag(p))

	io.WriteString(w, svgEnd)
}

func secondHandTag(p Point) string {
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;">%s</line>`, p.X, p.Y, SecondHandText)
}

func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)

	p.X = p.X * secondHandLength
	p.Y = p.Y * secondHandLength

	p.X = clockCentreX + p.X
	p.Y = clockCentreY - p.Y

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