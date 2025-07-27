package main

import (
	"fmt"
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
