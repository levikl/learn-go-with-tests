package clockface

import (
	"fmt"
	"io"
	"time"
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	clockCenterX     = 150
	clockCenterY     = 150
)

// SVGWriter writes an SVG representation of an analogue clock, showing the time t, to the writer w.
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart) // nolint:errcheck
	io.WriteString(w, bezel)    // nolint:errcheck
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd) // nolint:errcheck
}

func secondHand(w io.Writer, t time.Time) {
	p := makeHand(secondHandPoint(t), secondHandLength)
	fmt.Fprintf( // nolint:errcheck
		w,
		`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`,
		p.X,
		p.Y,
	)
}

func minuteHand(w io.Writer, t time.Time) {
	p := makeHand(minuteHandPoint(t), minuteHandLength)
	fmt.Fprintf( // nolint:errcheck
		w,
		`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`,
		p.X,
		p.Y,
	)
}

func hourHand(w io.Writer, t time.Time) {
	p := makeHand(hourHandPoint(t), hourHandLength)
	fmt.Fprintf( // nolint:errcheck
		w,
		`<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`,
		p.X,
		p.Y,
	)
}

func makeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}                // scale
	p = Point{p.X, -p.Y}                                 // flip
	return Point{p.X + clockCenterX, p.Y + clockCenterY} // translate
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
