package hexmap

import (
	"fmt"
	"math"

	"github.com/cshabsin/thegrid/js"
)

type gridEntry struct {
	x         int
	y         int
	meshShown bool
	centerX   float64
	centerY   float64
}

type HexMap struct {
	width     int
	height    int
	radius    float64
	staggerUp bool

	dx   float64
	dy   float64
	grid [][]gridEntry
}

var SinPiOver3 = math.Sin(math.Pi / 3)

func NewHexMap(width, height int, radius float64, staggerUp bool) *HexMap {
	hexmap := &HexMap{
		width:     width,
		height:    height,
		radius:    radius,
		staggerUp: staggerUp,
		dx:        radius * 1.5,
		dy:        radius * 2 * SinPiOver3,
	}
	grid := make([][]gridEntry, width)
	for col := range grid {
		grid[col] = make([]gridEntry, height)
		for row := range grid[col] {
			center := hexmap.calculateCenter(col, row)
			grid[col][row] = gridEntry{
				x:         col,
				y:         row,
				meshShown: true,
				centerX:   center.X,
				centerY:   center.Y,
			}
		}
	}
	hexmap.grid = grid
	return hexmap
}

// Returns true if the cell is a "down" cell in its row.
func (h HexMap) isCellDown(col, row int) bool {
	if h.staggerUp {
		return col%2 == 1
	}
	return col%2 == 0
}

func (h HexMap) calculateCenter(col, row int) js.SVGCoord {
	x := h.dx*float64(col) + h.radius
	y := h.dy*float64(row) + h.radius*SinPiOver3
	if h.isCellDown(col, row) {
		y += h.dy / 2
	}
	return js.SVGCoord{X: x, Y: y}
}

// GetPixHeight returns the height of the full hexmap in pixels.
func (h HexMap) GetPixHeight() float64 {
	return h.dy * (float64(h.height) + 0.5)
}

// getPixWidth returns the width of the full hexmap in pixels.
// This includes the > sticking out the side.
func (h HexMap) GetPixWidth() float64 {
	return h.dx*(float64(h.width)) + h.radius*math.Sin(math.Pi/6)
}

// GridMesh returns the SVG path for the grid starting with the top-left corner of the (0, 0) hex.
func (h HexMap) GridMesh() js.SVGPath {
	hexagon := Hexagon(h.radius)
	// drawn[0] = top left, going clockwise. 0, 1, 2, and 5 are
	// always true, while 3 and 4 are recalculated per cell to
	// avoid double-drawing any edges.
	drawn := []bool{true, true, true, true, true, true}
	var path js.SVGPath

	for col := 0; col < h.width; col++ {
		for row := 0; row < h.height; row++ {
			if !h.showMesh(col, row) {
				continue
			}
			path = path.MoveAbs(h.calculateCenter(col, row), false).MoveRel(hexagon[0], false)
			drawn[3] = h.isDownRightShown(col, row)
			drawn[4] = h.isDownShown(col, row)
			for i := 0; i < 6; i++ {
				path = path.MoveRel(hexagon[i+1], drawn[i])
			}
		}
	}
	return path
}

func (h HexMap) HexPath(document js.DOMDocument, cls string) js.DOMElement {
	hexagon := h.HexagonPath()
	return document.CreateSVG("path", js.MakeAttr("class", cls),
		js.MakeAttr("d", string(hexagon)),
	)
}

func (h HexMap) CellTranslate(col, row int) js.Attr {
	return js.MakeAttr("transform", fmt.Sprintf("translate(%f,%f)", h.grid[col][row].centerX, h.grid[col][row].centerY))
}

func (h HexMap) HexagonPath() js.SVGPath {
	var p js.SVGPath
	for i, coord := range Hexagon(h.radius) {
		if i == 0 {
			p = p + js.SVGPath(fmt.Sprintf("m%f,%f", coord.DX, coord.DY))
		} else {
			p = p + js.SVGPath(fmt.Sprintf("l%f,%f", coord.DX, coord.DY))
		}
	}
	return p
}

func (h HexMap) showMesh(col, row int) bool {
	return h.grid[col][row].meshShown
}

func (h HexMap) isDownRightShown(col, row int) bool {
	return true
}

func (h HexMap) isDownShown(col, row int) bool {
	return true
}

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
