package main

import (
	"fmt"
	"math"
	"path"

	"github.com/cshabsin/thegrid/example/server/data"
	"github.com/cshabsin/thegrid/example/view"
	"github.com/cshabsin/thegrid/hexmap"
	"github.com/cshabsin/thegrid/js"
	"github.com/cshabsin/thegrid/js/attr"
	"github.com/cshabsin/thegrid/js/svg"
	"github.com/cshabsin/thegrid/model"
)

func main() {
	url, err := js.URL()
	if err != nil {
		fmt.Println(err)
		return
	}
	document := js.Document()
	svgElem := svg.GetSVGById(document, "map-svg")
	if url.Fragment == "graph" {
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
			svgElem.Append(drawGraph(svgElem, elapsed))

			js.RequestAnimationFrame(animate)
		}
		js.RequestAnimationFrame(animate)

		select {}
	}
	hm := hexmap.NewHexMap(10, 11, 70, false)

	mapGroup := svgElem.CreateElement("g", attr.Make("class", "map-anchor-group"), attr.Make("transform", "translate(10,10)"))
	mapGroup.Append(hm.GridMesh().ToElement(svgElem, attr.Class("map-mesh")))

	newURL := *url
	newURL.Path = path.Join(newURL.Path, "/data")
	explorersSystemData, err := model.FromURL(newURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	mapView := &view.MapView{SVG: svgElem, HexMap: hm, DataElement: document.GetElementByID("data-contents")}
	for col := range explorersSystemData.HexGrid {
		for row := range explorersSystemData.HexGrid[col] {
			parsec := mapView.NewParsec(col, row, explorersSystemData.GetCell(col, row))
			mapGroup.Append(parsec.Anchor.AsDOM())
		}
	}

	for _, seg := range data.ExplorersPathData.Segments {
		mapGroup.Append(mapView.NewPathSegment(seg, "spiny-rat"))
	}

	svgElem.SetAttr("height", fmt.Sprintf("%fpx", hm.GetPixHeight()+20))
	svgElem.SetAttr("width", fmt.Sprintf("%fpx", hm.GetPixWidth()+20))
	svgElem.Append(mapGroup)

	select {}
}

func drawGraph(svgElem svg.SVG, timestamp float64) svg.Element {
	var g svg.G
	for i := 0; i < 10000; i += 70 {
		i64 := math.Mod(float64(i)+timestamp, 10000)
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
	return g.ToElement(svgElem)
}
