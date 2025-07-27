package main

import (
	"fmt"
	"math"

	"github.com/cshabsin/thegrid/js"
	"github.com/cshabsin/thegrid/js/attr"
	"github.com/cshabsin/thegrid/js/svg"
)

func main() {
	document := js.Document()
	svgElem := svg.GetSVGById(document, "map-svg")
	svgElem.SetAttr("height", "100%")
	svgElem.SetAttr("width", "100%")
	svgElem.SetAttr("viewBox", "0 0 10000 10000")

	var paths []*svg.Path
	for i := 0; i < 10000; i += 70 {
		path := &svg.Path{}
		path.SetAttr(attr.Make("style", fmt.Sprintf("fill: none; stroke: hsl(%d, 100%%, 50%%); stroke-width: 10px", i*360/10000)))
		svgElem.Append(path.ToElement(svgElem))
		paths = append(paths, path)
	}

	var startTime float64
	var animate func(timestamp float64)
	animate = func(timestamp float64) {
		if startTime == 0 {
			startTime = timestamp
		}
		elapsed := timestamp - startTime

		for i, path := range paths {
			i64 := math.Mod(float64(i*70)+elapsed*8.5, 10000)
			path.Reset()
			path.MoveAbs(svg.Coord{X: i64, Y: 0}, false)
			path.MoveAbs(svg.Coord{X: 0, Y: 10000 - i64}, true)

			path.MoveAbs(svg.Coord{X: 10000 - i64, Y: 10000}, false)
			path.MoveAbs(svg.Coord{X: 0, Y: 10000 - i64}, true)

			path.MoveAbs(svg.Coord{X: 10000 - i64, Y: 10000}, false)
			path.MoveAbs(svg.Coord{X: 10000, Y: i64}, true)

			path.MoveAbs(svg.Coord{X: i64, Y: 0}, false)
			path.MoveAbs(svg.Coord{X: 10000, Y: i64}, true)
			path.Update()
		}

		js.RequestAnimationFrame(animate)
	}
	js.RequestAnimationFrame(animate)

	select {}
}
