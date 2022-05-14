package main

import (
	"github.com/cshabsin/thegrid/hexmap"
	"github.com/cshabsin/thegrid/js"
)

func main() {
	document := js.Document()
	// p := document.CreateElement("p")
	// p.Set("innerHTML", "hi")
	// document.Body().Append(p)
	hm := hexmap.NewHexMap(10, 11, 70, false)
	mapGroup := document.CreateSVG("g", js.MakeAttr("class", "map-anchor-group"), js.MakeAttr("transform", "translate(10,10)"))
	mapGroup.Append(hm.GridMesh().MakeSVG(document, js.Class("map-mesh")))

	document.GetElementByID("map-contents").Append(
		document.CreateSVG("svg", js.MakeAttr("height", hm.GetPixHeight()+20), js.MakeAttr("width", hm.GetPixWidth()+20)).Append(mapGroup))
}
