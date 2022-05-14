package hexmap

import (
	"math"

	"github.com/cshabsin/thegrid/js"
)

// Returns a list of coordinates [x, y] representing a hexagon
// path starting with the left vertex, going clockwise, ending
// once again on the left vertex.
//
// The first coordinate returned is a relative movement from
// the center of the hex, the remaining coordinates are
// relative to the previous point.
//
// Example usage in an SVG path:
//     "m" + hexagon(30).join("l")
// draws a hexagon with the center of the hex set where the
// pointer originally started.

// Origin point for each relative coordinate.
func Hexagon(radius float64) []js.SVGVector {
	var x0, y0 float64
	var coords []js.SVGVector
	for i := 0; i < 7; i++ {
		angle := math.Pi * (0.5 + float64(i)/3)
		x1, y1 := math.Sin(angle)*radius, -math.Cos(angle)*radius
		coords = append(coords, js.SVGVector{DX: x1 - x0, DY: y1 - y0})
		x0, y0 = x1, y1
	}
	return coords
}
