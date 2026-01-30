package main

import (
	"fmt"
	"net/url"

	"github.com/cshabsin/thegrid/apps/explorers/data"
	"github.com/cshabsin/thegrid/apps/explorers/model"
	"github.com/cshabsin/thegrid/apps/explorers/view"
	"github.com/cshabsin/thegrid/hexmap"
	"github.com/cshabsin/thegrid/js"
	"github.com/cshabsin/thegrid/js/attr"
	"github.com/cshabsin/thegrid/js/svg"
)

func main() {
	document := js.Document()
	svgElem := svg.GetSVGById(document, "map-svg")
	hm := hexmap.NewHexMap(10, 11, 70, false)

	mapGroup := svgElem.CreateElement("g", attr.Make("class", "map-anchor-group"), attr.Make("transform", "translate(10,10)"))
	view.CreateStarfieldPattern(svgElem)

	newURL, err := url.Parse("data")
	if err != nil {
		fmt.Println(err)
		return
	}
	explorersSystemData, err := model.FromURL(newURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set the fill for each hex based on data.
	for col, colData := range explorersSystemData.HexGrid {
		for row := range colData {
			if explorersSystemData.GetCell(col, row) == nil {
				hm.Grid[col][row].Fill = "url(#starfield)"
			} else {
				hm.Grid[col][row].Fill = "#333"
			}
		}
	}

	// The order of appends here is important for layering.
	// 1. Fills for the hexes.
	mapGroup.Append(hm.CreateFillsGroup(svgElem))

	// 2. The grid mesh border.
	mapGroup.Append(hm.GridMesh().ToElement(svgElem, attr.Class("map-mesh")))

	// 3. The interactive elements (planets, labels, event listeners).
	highlighter := hm.HexagonPath().ToElement(svgElem, attr.ID("highlighter"), attr.Class("map-hexagon-hilite"), attr.Make("visibility", "hidden"), attr.Make("fill", "none"), attr.Make("pointer-events", "none"))
	selector := hm.HexagonPath().ToElement(svgElem, attr.ID("selector"), attr.Class("map-hexagon-selected"), attr.Make("visibility", "hidden"), attr.Make("fill", "none"), attr.Make("pointer-events", "none"))
	mapView := &view.MapView{SVG: svgElem, HexMap: hm, DataElement: document.GetElementByID("data-contents"), Highlighter: highlighter, Selector: selector}
	for col, colData := range explorersSystemData.HexGrid {
		for row := range colData {
			mapGroup.Append(mapView.CreateHexAnchor(col, row, explorersSystemData.GetCell(col, row)))
		}
	}

	// 4. The path segments.
	for _, seg := range data.ExplorersPathData.Segments {
		mapGroup.Append(mapView.NewPathSegment(seg, "spiny-rat"))
	}

	// 5. The highlighter and selector, which appear on top of everything.
	mapGroup.Append(selector)
	mapGroup.Append(highlighter)

	svgElem.SetAttr("height", fmt.Sprintf("%fpx", hm.GetPixHeight()+20))
	svgElem.SetAttr("width", fmt.Sprintf("%fpx", hm.GetPixWidth()+20))
	svgElem.Append(mapGroup)

	select {}
}
