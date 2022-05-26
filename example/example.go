package main

import (
	"fmt"
	"path"

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
	hm := hexmap.NewHexMap(10, 11, 70, false)

	svg := svg.GetSVGById(document, "map-svg")

	mapGroup := svg.CreateElement("g", attr.Make("class", "map-anchor-group"), attr.Make("transform", "translate(10,10)"))
	mapGroup.Append(hm.GridMesh().ToElement(svg, attr.Class("map-mesh")))

	newURL := *url
	newURL.Path = path.Join(newURL.Path, "/data")
	data, err := model.FromURL(newURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	mapView := view.NewMapView(svg, hm)
	for col := range data.HexGrid {
		for row := range data.HexGrid[col] {
			parsec := mapView.NewParsec(col, row, data.GetCell(col, row))
			mapGroup.Append(parsec.Anchor.AsDOM())
		}
	}

	svg.SetAttr("height", fmt.Sprintf("%fpx", hm.GetPixHeight()+20))
	svg.SetAttr("width", fmt.Sprintf("%fpx", hm.GetPixWidth()+20))
	svg.Append(mapGroup)

	select {}
}
