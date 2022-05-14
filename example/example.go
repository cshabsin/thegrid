package main

import (
	"github.com/cshabsin/thegrid/hexmap"
	"github.com/cshabsin/thegrid/js"
)

func main() {
	document := js.Document()
	p := document.CreateElement("p")
	p.Set("innerHTML", "hi")
	document.Body().Append(p)
	svg := document.CreateSVG()
	hm := hexmap.NewHexMap(2, 2, 70, false)
	svg.AddPath(hm.GridMesh())
	svg.SetAttr("height", hm.GetPixHeight())
	svg.SetAttr("width", hm.GetPixWidth())
	document.Body().Append(svg)
}
