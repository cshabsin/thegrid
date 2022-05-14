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
	svg := document.CreateSVG()
	hm := hexmap.NewHexMap(2, 2, 70, false)
	svg.AddPath(hm.GridMesh(), js.Class("map-mesh"))
	svg.SetAttr("height", hm.GetPixHeight()+10)
	svg.SetAttr("width", hm.GetPixWidth()+10)
	document.Body().Append(svg)
}
