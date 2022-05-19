package main

import (
	"fmt"

	"github.com/cshabsin/thegrid/hexmap"
	"github.com/cshabsin/thegrid/js"
)

func main() {
	document := js.Document()
	hm := hexmap.NewHexMap(10, 11, 70, false)
	mapGroup := document.CreateSVG("g", js.MakeAttr("class", "map-anchor-group"), js.MakeAttr("transform", "translate(10,10)"))
	mapGroup.Append(hm.GridMesh().MakeSVG(document, js.Class("map-mesh")))
	mapGroup.Append(hm.HexPath(document, 3, 3, "map-hexagon-hilite"))

	svg := document.GetElementByID("map-svg")
	svg.SetAttr("height", fmt.Sprintf("%fpx", hm.GetPixHeight()+20))
	svg.SetAttr("width", fmt.Sprintf("%fpx", hm.GetPixWidth()+20))
	svg.Append(mapGroup)
}
