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

	var startTime float64
	var animate func(timestamp float64)
	animate = func(timestamp float64) {
		if startTime == 0 {
			startTime = timestamp
		}
		elapsed := timestamp - startTime

		svgElem.Clear()
		svgElem.Append(drawGraph(elapsed).ToElement(svgElem))

		js.RequestAnimationFrame(animate)
	}
	js.RequestAnimationFrame(animate)

	select {}
}

func drawGraph(timestamp float64) *svg.G {
	var g svg.G
	for i := 0; i < 10000; i += 70 {
		i64 := math.Mod(float64(i)+timestamp*8.5, 10000)
		var p svg.Path
		p.MoveAbs(svg.Coord{X: i64, Y: 0}, false)
		p.MoveAbs(svg.Coord{X: 0, Y: 10000 - i64}, true)

		p.MoveAbs(svg.Coord{X: 10000 - i64, Y: 10000}, false)
		p.MoveAbs(svg.Coord{X: 0, Y: 10000 - i64}, true)

		p.MoveAbs(svg.Coord{X: 10000 - i64, Y: 10000}, false)
		p.MoveAbs(svg.Coord{X: 10000, Y: i64}, true)

		p.MoveAbs(svg.Coord{X: i64, Y: 0}, false)
		p.MoveAbs(svg.Coord{X: 10000, Y: i64}, true)

		p.SetAttr(attr.Make("style", fmt.Sprintf("fill: none; stroke: hsl(%d, 100%%, 50%%); stroke-width: 10px", i*360/10000)))
		g = g.Append(&p)
	}
	return &g
}
